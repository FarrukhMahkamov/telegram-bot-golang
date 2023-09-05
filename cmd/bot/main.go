package main

import (
	"log"

	"github.com/FarrukhMahkamov/telegram-bot-golang/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6302249545:AAEVikpNbfLScAZdnHnfn7Cw1EYZGuliSx8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
