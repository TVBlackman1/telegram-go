package handlers

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type StatesHandler struct {
	stateService *service.StateService
}

func NewStatesHandler(stateService *service.StateService) *StatesHandler {
	return &StatesHandler{stateService}
}

func (handler *StatesHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	currentState, _ := handler.stateService.GetCurrentState(message)
	return types.MessageUnion{
		Text: fmt.Sprintf("Your state is %s", currentState.Name),
	}
}
