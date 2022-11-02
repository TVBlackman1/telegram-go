package telegramlistener

import (
	"log"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/router"
	"github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgWorkspace struct {
	token    string
	router   *router.Router
	notifier *notifier.SystemNotifier
	bot      *tgbotapi.BotAPI
}

func NewTelegramBot(token string, router *router.Router, notifier *notifier.SystemNotifier) *TgWorkspace {
	return &TgWorkspace{
		token:    token,
		router:   router,
		bot:      nil,
		notifier: notifier,
	}
}

func (workspace *TgWorkspace) Run() error {
	bot, err := tgbotapi.NewBotAPI(workspace.token)
	if err != nil {
		return err
	}
	workspace.bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)

	go workspace.RunNotificator()
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
	answersContent := usingHandler.Process(receivedMessage)
	chatId := receivedMessage.Sender.ChatId
	workspace.sendMessages(chatId, &answersContent)
}

func (workspace *TgWorkspace) sendAutoMessage(chatId types.ChatId) {
	systemHandler := workspace.router.GetSystemMessager()
	answersContent := systemHandler.Process(chatId)
	workspace.sendMessages(chatId, &answersContent)
}

func (workspace *TgWorkspace) SendAutoMessageWithContext(chatId types.ChatId, context interface{}) {
	panic("not implemented")
}

func (workspace *TgWorkspace) RunNotificator() {
	notificator := workspace.notifier.GetNotificator()
	for notification := range notificator {
		chatId := notification.ChatId
		workspace.sendAutoMessage(chatId)
	}
}

func (workspace *TgWorkspace) sendMessages(chatId types.ChatId, messagesContent *handlers.HandlerProcessResult) {
	for _, answer := range messagesContent.Messages {
		msg := tgbotapi.NewMessage(int64(chatId), "")
		presenter.Present(&msg, answer)
		workspace.bot.Send(msg) // TODO add delay with condition
	}
	for _, automessage := range messagesContent.Automessages {
		// TODO global change, with timers, etc
		workspace.sendAutoMessage(automessage.ChatId)
	}
}

func (workspace *TgWorkspace) buildReceivedMessage(message *tgbotapi.Message) types.ReceivedMessage {
	chatId := types.ChatId(message.Chat.ID)
	sender := types.Sender{
		ChatId: chatId,
		Name:   message.From.FirstName,
		Login:  message.From.UserName,
	}
	return types.ReceivedMessage{
		Sender:  sender,
		Content: presenter.Collect(message),
	}
}

func (workspace *TgWorkspace) getMessageChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return workspace.bot.GetUpdatesChan(u)
}
