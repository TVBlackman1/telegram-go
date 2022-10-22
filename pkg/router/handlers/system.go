package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type SystemHandler struct {
	userService *service.UserService
}

func NewSystemHandler(stateService *service.UserService) *SystemHandler {
	return &SystemHandler{stateService}
}

// gets command from enum to exec something
func (handler *SystemHandler) Process(chatId types.ChatId) HandlerProcessResult {
	currentState, _ := handler.userService.GetCurrentState(chatId)
	stateProcessor, _ := handler.userService.GetCurrentStateProcessor(currentState)
	// stateProcessor.SetContext(, currentState.Id)
	stateProcessor.ProcessSystemInvoke(chatId)
	// TODO several times sends message to processor
	return HandlerProcessResult{
		Messages:      stateProcessor.GetBotMessages(),
		Notifications: stateProcessor.GetNotifications(),
	}
}
