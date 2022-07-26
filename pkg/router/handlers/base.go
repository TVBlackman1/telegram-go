package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

type HandlerProcessResult struct {
	Messages     []types.Message
	Automessages []notifier.NotifierContext
}

type ConcreteHandler interface {
	Process(message types.ReceivedMessage) HandlerProcessResult
}
