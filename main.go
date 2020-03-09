package main

import (
	// "fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := processUpdate(update)

		bot.Send(msg)
	}
}

func processUpdate(update tgbotapi.Update) tgbotapi.Chattable {
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	replyText := replyMessageText(update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
	msg.ReplyToMessageID = update.Message.MessageID

	if update.Message.Text == "/start" {
		msg.ReplyMarkup = numericKeyboard
	}
	return msg
}
