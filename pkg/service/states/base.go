package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

type UserState interface {
	PreparePresentation() types.MessageUnion
	ProcessUserInput(msg types.ReceivedMessage)
	SetContext(msg types.ReceivedMessage, context interface{}) error
	GetName() string
}
