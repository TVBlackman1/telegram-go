package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

type UserState interface {
	// PreparePresentation() types.MessageUnion
	ProcessUserInput(msg types.ReceivedMessage)
	ProcessSystemInvoke(chatId types.ChatId)
	GetBotMessages() []types.MessageUnion
	GetNotifications() []notifier.NotifierContext
	ProcessContextedSystemInvoke(chatId types.ChatId, context interface{})
	SetContext(msg types.ReceivedMessage, context interface{}) error
	GetName() string
}
