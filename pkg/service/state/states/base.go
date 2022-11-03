package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

type UserState interface {
	ProcessUserInput(msg types.ReceivedMessage)
	ProcessSystemInvoke(chatId types.ChatId)
	GetBotMessages() []types.Message
	GetAutoMessages() []notifier.NotifierContext
	ProcessContextedSystemInvoke(chatId types.ChatId, context interface{})
	SetState(msg types.ReceivedMessage, state repository.StateDbDto) error
	GetName() string
}
