package impls

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

type ConcreteHandler interface {
	Process(message types.ReceivedMessage) types.MessageUnion
}
