package discorde

import "github.com/ldtrieu/cerberus/package/discord"

func Init(notifyApi discord.IDiscordNotification, options *ClientOptions) error {
	hub := CurrentHub()
	client, err := NewClient(notifyApi, options)
	if err != nil {
		return err
	}
	hub.BindClient(client)
	return nil
}

func CaptureExeption(exception error) {
	hub := CurrentHub()
	hub.CaptureException(exception)
}

func CaptureMessage(message string) {
	hub := CurrentHub()
	hub.CaptureMessage(message)
}

func WithScope(f func(scope *Scope)) {
	hub := CurrentHub()
	hub.WithScope(f)
}
