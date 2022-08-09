package handler

import (
	handlers "github.com/TVBlackman1/telegram-go/pkg/handler/impls"
	"github.com/TVBlackman1/telegram-go/pkg/states"
)

type Handler struct {
	StartListener  *handlers.StartHandler
	StatesListener *handlers.StatesHandler
	TestListener   *handlers.TestHandler
}

func NewHandler(stateService *states.StateService) *Handler {
	return &Handler{
		StartListener:  handlers.NewStartHandler(stateService),
		StatesListener: handlers.NewStatesHandler(stateService),
		TestListener:   handlers.NewTestHandler(stateService),
	}
}
