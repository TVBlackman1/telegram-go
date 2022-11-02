package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type HelpHandler struct {
	stateService *service.UserService
}

func NewHelpHandler(stateService *service.UserService) *HelpHandler {
	return &HelpHandler{stateService}
}

func (handler *HelpHandler) Process(message types.ReceivedMessage) HandlerProcessResult {
	return HandlerProcessResult{
		Messages: []types.Message{
			{Text: "Just test it and be happy :)"},
		},
	}
}
