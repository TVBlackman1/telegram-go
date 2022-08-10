package handlers

import (
	"fmt"

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
	fmt.Printf("Start from chat id %d\n", message.ChatId)
	// state := handler.stateService.GetCurrentState(message)
	// TODO add check current state, nil
	handler.stateService.RegisterNewUser(message.ChatId)
	return types.MessageUnion{
		Text: fmt.Sprintf("New user with chat id %d", message.ChatId),
	}
}
