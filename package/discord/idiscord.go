package discord

type IDiscordNotification interface {
	SendError(message string, fields ...EmbedFields) error
	SendNotification(color string, message string, fields ...EmbedFields) error
}
