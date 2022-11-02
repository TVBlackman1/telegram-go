package router

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	handlers "github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type routerHandlers struct {
	StartHandler *handlers.StartHandler
	UserHandler  *handlers.UserHandler
	TestHandler  *handlers.TestHandler
}
type Router struct {
	handlers       routerHandlers
	systemMessager *SystemMessager
}

func NewRouter(userService *service.UserService) *Router {
	return &Router{
		handlers: routerHandlers{
			StartHandler: handlers.NewStartHandler(userService),
			UserHandler:  handlers.NewUserHandler(userService),
			TestHandler:  handlers.NewTestHandler(userService),
		},
		systemMessager: NewSystemMessager(userService),
	}
}

func (router *Router) RouteByMessage(message types.ReceivedMessage) handlers.ConcreteHandler {
	var retHandler handlers.ConcreteHandler
	// TODO add more handlers: about, help
	if message.Content.Text == "/start" {
		retHandler = router.handlers.StartHandler
	} else {
		retHandler = router.handlers.UserHandler
	}
	return retHandler
}

// TODO think about output type
func (router *Router) GetSystemMessager() *SystemMessager {
	return router.systemMessager
}
