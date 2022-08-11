package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type StartHandler struct {
	stateService *service.StateService
}

func NewStartHandler(stateService *service.StateService) *StartHandler {
	return &StartHandler{stateService}
}

func (handler *StartHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	return handler.stateService.RegisterNewUser(message.Sender)
}
