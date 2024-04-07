package main

import (
	"log"

	myTelegramBot "github.com/Sanchir01/bot_crypto/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7087783787:AAGmB1OViUpZUX4X5WQd0Wn1mmbYK1xyGVE")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	telegramBot := myTelegramBot.NewBot(bot)

	telegramBot.Start()
}
