package main

import (
	"log"

	"github.com/TVBlackman1/telegram-go/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config, err := configs.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_TOKEN)
	if err != nil {
		log.Panic("telegram-bot token exception: ", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}
