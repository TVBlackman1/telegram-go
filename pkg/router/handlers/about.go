package handlers

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type AboutHandler struct {
	stateService *service.UserService
}

func NewAboutHandler(stateService *service.UserService) *AboutHandler {
	return &AboutHandler{stateService}
}

func (handler *AboutHandler) Process(message types.ReceivedMessage) HandlerProcessResult {
	return HandlerProcessResult{
		Messages: []types.Message{
			{Text: "Bot template of tvblackman1. See more:\nhttps://github.com/TVBlackman1/telegram-go"},
		},
	}
}
