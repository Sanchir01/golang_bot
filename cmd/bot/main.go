package main

import (
	"log"
	"os"

	myTelegramBot "github.com/Sanchir01/bot_crypto/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient(os.Getenv("POCKET_CONSUMER_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := myTelegramBot.NewBot(bot, pocketClient, "http://localhost/")

	telegramBot.Start()
}
