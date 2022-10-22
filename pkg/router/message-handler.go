package router

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	handlers "github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type Router struct {
	StartHandler  *handlers.StartHandler
	UserHandler   *handlers.UserHandler
	SystemHandler *handlers.SystemHandler
	TestHandler   *handlers.TestHandler
}

func NewRouter(userService *service.UserService) *Router {
	return &Router{
		StartHandler:  handlers.NewStartHandler(userService),
		UserHandler:   handlers.NewUserHandler(userService),
		SystemHandler: handlers.NewSystemHandler(userService),
		TestHandler:   handlers.NewTestHandler(userService),
	}
}

func (handler *Router) RouteByMessage(message types.ReceivedMessage) handlers.ConcreteHandler {
	var retHandler handlers.ConcreteHandler
	if message.Content.Text == "/start" {
		retHandler = handler.StartHandler
	} else {
		retHandler = handler.UserHandler
	}
	return retHandler
}

// TODO think about output type
func (handler *Router) GetSystemHandler() *handlers.SystemHandler {
	return handler.SystemHandler
}
