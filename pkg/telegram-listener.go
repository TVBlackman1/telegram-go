package pkg

import (
	"log"

	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	"github.com/TVBlackman1/telegram-go/pkg/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunTelegramBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
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

			state := states.DefineState()
			userInput := presenter.MessageUnion{
				Text: update.Message.Text,
			}
			state.ProcessUserInput(update.Message.From, userInput)
			answer := state.PreparePresentation()

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			presenter.Present(&msg, answer)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}
