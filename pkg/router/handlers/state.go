package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
	"github.com/TVBlackman1/telegram-go/pkg/service/states"
)

type StatesHandler struct {
	stateService *service.StateService
}

func NewStatesHandler(stateService *service.StateService) *StatesHandler {
	return &StatesHandler{stateService}
}

func (handler *StatesHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	currentState, _ := handler.stateService.GetCurrentState(message)
	stateProcessor := states.DefineState(currentState.Name)
	stateProcessor.SetContext(message, currentState.Id)
	stateProcessor.ProcessUserInput(message)
	return stateProcessor.PreparePresentation()
}
