package discorde

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/ldtrieu/cerberus/package/discord"
)

var (
	hostName, _ = os.Hostname()
	greenColor  = "#00FF00"
)

type Client struct {
	options   *ClientOptions
	notifyApi discord.IDiscordNotification
}

func NewClient(notifyApi discord.IDiscordNotification, options *ClientOptions) (*Client, error) {
	if notifyApi == nil {
		return nil, errors.New("notification api cant be nil")
	}

	if options.ProjectName == "" {
		return nil, errors.New("project name is required")
	}

	if options.Environment == "" {
		return nil, errors.New("environment is required")
	}

	ops := options
	ops.MaxErrorDepth = 10

	return &Client{
		options:   ops,
		notifyApi: notifyApi,
	}, nil
}

func (c *Client) eventFromException(exception error) *Event {
	err := exception
	if err == nil {
		err = fmt.Errorf("%s call with nil error", callerFunctionName())
	}

	event := &Event{}
	event.Exception = &Exception{
		Value:      err.Error(),
		Type:       reflect.TypeOf(err).String(),
		Stacktrace: ExtractStacktrace(err),
	}

	if event.Exception.Stacktrace == nil {
		event.Exception.Stacktrace = NewStacktrace()
	}

	var title strings.Builder
	title.WriteString(fmt.Sprintf("%s - %s\n", event.Exception.Type, event.Exception.Value))
	if event.Exception.Stacktrace != nil {
		for _, f := range event.Exception.Stacktrace.Frames {
			title.WriteString(fmt.Sprintf("%s:%d %s(0x%x)\n", f.Filename, f.Lineno, f.Function, f.ProgramCounter))
		}
	}

	event.Message = title.String()

	return event
}

func (c *Client) prepareEvent(event *Event, scope *Scope) *Event {
	event.Timestamp = time.Now()

	event.ServerName = hostName
	event.Platform = "go"
	event.Environment = c.options.Environment
	event.ProjectName = c.options.ProjectName
	event.Release = getReleaseVersion()
	event.Arch = runtime.GOARCH
	event.NumCPU = runtime.NumCPU()
	event.GOOS = runtime.GOOS
	event.GoVersion = runtime.Version()

	if scope != nil {
		event = scope.ApplyToEvent(event)
	}

	return event
}

func extractCommonFields(event *Event) []discord.EmbedFields {
	var fields []discord.EmbedFields
	fields = append(fields, discord.EmbedFields{
		Name:  "Project",
		Value: event.ProjectName,
	}, discord.EmbedFields{
		Name:  "Runtime",
		Value: event.GoVersion,
	}, discord.EmbedFields{
		Name:  "Server",
		Value: event.ServerName,
	})

	return fields
}

func (c *Client) sendMessageException(event *Event) error {
	if c.notifyApi == nil {
		return errors.New("notifyApi is nil")
	}

	fields := extractCommonFields(event)
	for k, v := range event.Tags {
		fields = append(fields, discord.EmbedFields{
			Name:  k,
			Value: v,
		})
	}

	return c.notifyApi.SendError(event.Message, fields...)
}

func (c *Client) sendMessage(event *Event) error {
	if c.notifyApi == nil {
		return errors.New("notifyApi is nil")
	}

	fields := extractCommonFields(event)
	for k, v := range event.Tags {
		fields = append(fields, discord.EmbedFields{
			Name:  k,
			Value: v,
		})
	}

	return c.notifyApi.SendNotification(greenColor, event.Message, fields...)
}

func (c *Client) CaptureException(exception error, scope *Scope) {
	event := c.eventFromException(exception)
	event = c.prepareEvent(event, scope)
	c.sendMessageException(event)

}

func (c *Client) CaptureMessage(message string, scope *Scope) {
	event := c.prepareEvent(&Event{Message: message}, scope)
	c.sendMessage(event)
}
