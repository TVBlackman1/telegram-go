package telegramlistener

import (
	"log"
	"reflect"

	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	token   string
	handler *Handler
}

func NewTelegramBot(token string, handler *Handler) *TelegramBot {
	return &TelegramBot{
		token, handler,
	}
}

func (telegramBot *TelegramBot) Run() error {
	bot, err := tgbotapi.NewBotAPI(telegramBot.token)
	if err != nil {
		return err
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			answer := telegramBot.allocateMessage(update.Message)
			if !reflect.ValueOf(answer).IsZero() {
				presenter.Present(&msg, answer)
				bot.Send(msg)
			}

		}
	}
	return nil

}

func (telegramBot *TelegramBot) allocateMessage(message *tgbotapi.Message) presenter.MessageUnion {
	// TODO add interfaces to listeners
	var retMessage presenter.MessageUnion
	if message.Text == "/start" {
		telegramBot.handler.StartListener.Process(message)
	} else {
		retMessage = telegramBot.handler.TestListener.Process(message)
	}
	return retMessage
}
