package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(stateService *service.UserService) *UserHandler {
	return &UserHandler{stateService}
}

func (handler *UserHandler) Process(message types.ReceivedMessage) HandlerProcessResult {
	chatId := message.Sender.ChatId
	currentState, _ := handler.userService.GetCurrentState(chatId)
	stateProcessor, _ := handler.userService.GetCurrentStateProcessor(currentState)
	stateProcessor.SetContext(message, currentState.Id)
	// TODO change ID to full state content in above method
	stateProcessor.ProcessUserInput(message)
	// TODO several times sends message to processor
	return HandlerProcessResult{
		Messages:      stateProcessor.GetBotMessages(),
		Notifications: stateProcessor.GetNotifications(),
	}
}
