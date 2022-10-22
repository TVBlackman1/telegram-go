package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type StartHandler struct {
	stateService *service.UserService
}

func NewStartHandler(stateService *service.UserService) *StartHandler {
	return &StartHandler{stateService}
}

func (handler *StartHandler) Process(message types.ReceivedMessage) HandlerProcessResult {
	return HandlerProcessResult{
		Messages: []types.MessageUnion{
			handler.stateService.RegisterNewUser(message.Sender),
		},
	}
}
