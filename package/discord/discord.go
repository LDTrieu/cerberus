package discord

import (
	"fmt"
	"strings"
	"time"
)

const titleLength = 100

type client struct {
	webhookURL string
	user       string
	avatar     string
}

func Init(url, user, avatar string) IDiscordNotification {
	return &client{
		webhookURL: url,
		user:       user,
		avatar:     avatar,
	}
}

func (c *client) SendError(message string, fields ...EmbedFields) error {
	//set webhook
	wh := Webhook{Username: c.user, AvatarUrl: c.avatar}

	title := message
	if len(message) > titleLength {
		title = fmt.Sprintf("%s ...", strings.TrimRight(message, message[titleLength-1:]))
	}

	//init embed
	emb := Embed{
		Title:       title,
		Description: message,
		Url:         "https://go.dev/doc/modules/publishing",
		Timestamp:   time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:       getColor("#FF0000"),
		Footer: EmbedFooter{
			Text:    "Sent via github.com/ldtrieu",
			IconUrl: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
		Thumbnail: EmbedThumbnail{
			Url: "https://cdn.iconscout.com/icon/free/png-256/elastic-283142.png", //"https://i.ibb.co/mDjnNPT/download.png",
		},
		Author: EmbedAuthor{
			Name:    "WsBlue",
			Url:     "https://github.com/ldtrieu",
			IconUrl: "https://e7.pngegg.com/pngimages/828/497/png-clipart-computer-icons-github-bitbucket-software-repository-cosmic-blue-angle.png",
		},
	}
	emb.Fields = append(emb.Fields, fields...)

	wh.Embeds = append(wh.Embeds, emb)

	return sendMessage(c.webhookURL, wh, true)
}

func (c *client) SendNotification(color string, message string, fields ...EmbedFields) error {
	wh := Webhook{
		Username:  c.user,
		AvatarUrl: c.avatar,
	}

	title := message
	if len(message) > titleLength {
		title = fmt.Sprintf("%s ...", strings.TrimRight(message, message[titleLength-1:]))
	}

	//init embed
	emb := Embed{
		Title:       title,
		Description: message,
		Url:         "https://go.dev/doc/modules/publishing",
		Timestamp:   time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:       getColor(color),
		Footer: EmbedFooter{
			Text:    "Sent via github.com/ldtrieu",
			IconUrl: "https://i.ibb.co/mDjnNPT/download.png",
		},
		Thumbnail: EmbedThumbnail{
			Url: "https://cdn.iconscout.com/icon/free/png-256/elastic-283142.png",
		},
		Author: EmbedAuthor{
			Name:    "WsBlue",
			Url:     "https://github.com/ldtrieu",
			IconUrl: "https://e7.pngegg.com/pngimages/828/497/png-clipart-computer-icons-github-bitbucket-software-repository-cosmic-blue-angle.png",
		},
	}
	emb.Fields = append(emb.Fields, fields...)

	wh.Embeds = append(wh.Embeds, emb)

	return sendMessage(c.webhookURL, wh, true)
}
