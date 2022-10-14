package router

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	handlers "github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type Router struct {
	StartHandler  *handlers.StartHandler
	StatesHandler *handlers.StatesHandler
	TestHandler   *handlers.TestHandler
}

func NewRouter(stateService *service.StateService) *Router {
	return &Router{
		StartHandler:  handlers.NewStartHandler(stateService),
		StatesHandler: handlers.NewStatesHandler(stateService),
		TestHandler:   handlers.NewTestHandler(stateService),
	}
}

func (handler *Router) RouteByMessage(message types.ReceivedMessage) handlers.ConcreteHandler {
	var retHandler handlers.ConcreteHandler
	if message.Content.Text == "/start" {
		retHandler = handler.StartHandler
	} else {
		retHandler = handler.StatesHandler
	}
	return retHandler
}
