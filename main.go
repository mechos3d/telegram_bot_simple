package main

import (
	// "fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
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

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		replyText := processMessage(update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
