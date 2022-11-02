package router

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	handlers "github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type routerHandlers struct {
	StartHandler *handlers.StartHandler
	UserHandler  *handlers.UserHandler
	AboutHandler *handlers.AboutHandler
	HelpHandler  *handlers.HelpHandler
}
type Router struct {
	handlers       routerHandlers
	systemMessager *SystemMessager
}

func NewRouter(userService *service.UserService) *Router {
	return &Router{
		handlers: routerHandlers{
			UserHandler:  handlers.NewUserHandler(userService),
			StartHandler: handlers.NewStartHandler(userService),
			AboutHandler: handlers.NewAboutHandler(userService),
			HelpHandler:  handlers.NewHelpHandler(userService),
		},
		systemMessager: NewSystemMessager(userService),
	}
}

func (router *Router) RouteByMessage(message types.ReceivedMessage) handlers.ConcreteHandler {
	var retHandler handlers.ConcreteHandler
	if message.Content.Text == "/start" {
		retHandler = router.handlers.StartHandler
	} else if message.Content.Text == "/about" {
		retHandler = router.handlers.AboutHandler
	} else if message.Content.Text == "/help" {
		retHandler = router.handlers.HelpHandler
	} else {
		retHandler = router.handlers.UserHandler
	}
	return retHandler
}

func (router *Router) GetSystemMessager() *SystemMessager {
	return router.systemMessager
}
