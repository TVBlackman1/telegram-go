package impls

import "github.com/TVBlackman1/telegram-go/pkg/service"

type StatesHandler struct {
	stateService *service.StateService
}

func NewStatesHandler(stateService *service.StateService) *StatesHandler {
	return &StatesHandler{stateService}
}
