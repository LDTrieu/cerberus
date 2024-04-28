package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getColor(hexColor string) int {
	hexColor = strings.Replace(hexColor, "#", "", -1)
	decimalColor, err := strconv.ParseInt(hexColor, 16, 64)
	if err != nil {
		return 0
	}
	return int(decimalColor)
}

func sendMessage(webookUrl string, content Webhook, retryOnRateLimit bool) error {
	if content.Content == "" && len(content.Embeds) == 0 {
		return errors.New("you must attach at least one of these: content; embeds")
	}
	if len(content.Embeds) > 10 {
		return errors.New("maximum number of embeds per webhook is 10")
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return err
	}
	for {
		res, err := http.Post(webookUrl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		if _, err = io.ReadAll(res.Body); err != nil {
			return err
		}
		res.Body.Close()

		switch res.StatusCode {
		case 204:
			return nil
		case 429:
			return fmt.Errorf("to many request(status code %d)", res.StatusCode)
		default:
			return fmt.Errorf("bad request (status code %d)", res.StatusCode)
		}
	}
}
