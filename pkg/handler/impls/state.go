package impls

import "github.com/TVBlackman1/telegram-go/pkg/states"

type StatesHandler struct {
	stateService *states.StateService
}

func NewStatesHandler(stateService *states.StateService) *StatesHandler {
	return &StatesHandler{stateService}
}
