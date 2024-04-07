package myTelegramBot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	replyStart   = "Hello, тебе надо дать мне доступ для сохранения ссылок %s"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {

	case commandStart:
		return b.handleStartCommand(message)
	default:

		return b.handleUnknownCommand(message)
	}

}
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	authorization, err := b.generateAuthorization(message.Chat.ID)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(replyStart, authorization))
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, Такой команды нету")
	_, err := b.bot.Send(msg)
	return err
}
