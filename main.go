package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Message from: %s: %s", update.Message.From, update.Message.Text)

		if update.Message.Command() == "cotacao" {
			getQuotations(bot, update)
		}
	}
}

func getQuotations(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	client := NewQuotationClient()
	quotations := client.GetQuotations()
	message := Formatter{quotations}.Format()

	log.Printf("Message [%s]", message)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
