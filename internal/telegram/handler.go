package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function handles a command from a user
func (b *Bot) handleCommand(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	msg.ReplyToMessageID = message.MessageID
	b.bot.Send(msg)
}

// This function handles a message from a user
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	msg.ReplyToMessageID = message.MessageID
	b.bot.Send(msg)
}
