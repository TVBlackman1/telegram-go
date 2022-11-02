package notifier

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

type NotifierParams struct {
	Time string
}

type NotifierContext struct {
	ChatId  types.ChatId
	Message types.Message
	Params  NotifierParams
}

type SystemNotifier struct {
	notifies chan NotifierContext
}

func NewSystemNotifier() *SystemNotifier {
	notifies := make(chan NotifierContext)
	return &SystemNotifier{notifies}
}

func (notifier *SystemNotifier) GetNotificator() chan NotifierContext {
	return notifier.notifies
}

func (notifier *SystemNotifier) Restart() {
	panic("not implemented")
}

func (notifier *SystemNotifier) NotifyInChat(chatId types.ChatId, params NotifierParams) {
	panic("not implemented")
}

func (notifier *SystemNotifier) NotifyInChatWithContext(chatId types.ChatId, context interface{}, params NotifierParams) {
	panic("not implemented")
}
