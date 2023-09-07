package main

import (
	"log"

	"github.com/FarrukhMahkamov/telegram-bot-golang/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6302249545:AAEVikpNbfLScAZdnHnfn7Cw1EYZGuliSx8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("108795-2f7f41f145cce3baf7403a2")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "https://localhost/")

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
