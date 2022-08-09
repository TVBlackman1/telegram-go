package telegramlistener

import (
	"log"
	"reflect"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgWorkspace struct {
	token  string
	router *router.Router
	bot    *tgbotapi.BotAPI
}

func NewTelegramBot(token string, router *router.Router) *TgWorkspace {
	return &TgWorkspace{
		token, router, nil,
	}
}

func (workspace *TgWorkspace) Run() error {
	bot, err := tgbotapi.NewBotAPI(workspace.token)
	if err != nil {
		return err
	}
	workspace.bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)
	updates := workspace.getMessageChan()

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			workspace.reactOnMessage(update.Message)
		}
	}
	return nil
}

func (workspace *TgWorkspace) reactOnMessage(message *tgbotapi.Message) {
	receivedMessage := workspace.buildReceivedMessage(message)
	usingHandler := workspace.router.RouteByMessage(receivedMessage)
	answer := usingHandler.Process(receivedMessage)
	if reflect.ValueOf(answer).IsZero() {
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	presenter.Present(&msg, answer)
	workspace.bot.Send(msg)
}

func (workspace *TgWorkspace) buildReceivedMessage(message *tgbotapi.Message) types.ReceivedMessage {
	chatId := types.ChatId(message.Chat.ID)
	return types.ReceivedMessage{
		ChatId:  chatId,
		Content: presenter.Collect(message),
	}
}

func (workspace *TgWorkspace) getMessageChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return workspace.bot.GetUpdatesChan(u)
}
