package router

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	handlers "github.com/TVBlackman1/telegram-go/pkg/router/handlers"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type SystemMessager struct {
	userService *service.UserService
}

func NewSystemMessager(stateService *service.UserService) *SystemMessager {
	return &SystemMessager{stateService}
}

func (messager *SystemMessager) Process(chatId types.ChatId) handlers.HandlerProcessResult {
	currentState, _ := messager.userService.GetCurrentState(chatId)
	stateProcessor, _ := messager.userService.GetCurrentStateProcessor(currentState)
	// stateProcessor.SetContext(, currentState.Id)
	stateProcessor.ProcessSystemInvoke(chatId)
	return handlers.HandlerProcessResult{
		Messages:     stateProcessor.GetBotMessages(),
		Automessages: stateProcessor.GetAutoMessages(),
	}
}

func (messager *SystemMessager) ProcessWithContext(chatId types.ChatId) handlers.HandlerProcessResult {
	currentState, _ := messager.userService.GetCurrentState(chatId)
	stateProcessor, _ := messager.userService.GetCurrentStateProcessor(currentState)
	// stateProcessor.SetContext(, currentState.Id)
	stateProcessor.ProcessSystemInvoke(chatId)
	return handlers.HandlerProcessResult{
		Messages:     stateProcessor.GetBotMessages(),
		Automessages: stateProcessor.GetAutoMessages(),
	}
}
