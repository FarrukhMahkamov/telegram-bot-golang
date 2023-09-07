package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	comandStart = "start"

	replyStartTemplate = "Hi. For saving links to your pocket account, you should give permission to me. Click this likn: \n%s"
)

// This function handles a command from a user
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case comandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

// This function handles a message from a user
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	msg.ReplyToMessageID = message.MessageID
	b.bot.Send(msg)
}

// This function handles a /start command
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(replyStartTemplate, authLink))
	_, err = b.bot.Send(msg)
	return err
}

// This function handles a unknown commands and returns defaul message
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I dont know this command 😕")

	_, err := b.bot.Send(msg)
	return err
}
