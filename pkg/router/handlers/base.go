package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

type HandlerProcessResult struct {
	Messages      []types.MessageUnion
	Notifications []notifier.NotifierContext
}

type ConcreteHandler interface {
	Process(message types.ReceivedMessage) HandlerProcessResult
}
