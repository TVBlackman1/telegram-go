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
	// TODO check user exists
	chatId := message.Sender.ChatId
	currentState, _ := handler.userService.GetCurrentState(chatId)
	stateProcessor, _ := handler.userService.GetCurrentStateProcessor(currentState)
	stateProcessor.SetState(message, currentState)
	// TODO change ID to full state content in above method
	stateProcessor.ProcessUserInput(message)
	return HandlerProcessResult{
		Messages:     stateProcessor.GetBotMessages(),
		Automessages: stateProcessor.GetAutoMessages(),
	}
}
