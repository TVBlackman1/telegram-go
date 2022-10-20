package handlers

import (
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type TestHandler struct {
	stateService *service.UserService
}

func NewTestHandler(stateService *service.UserService) *TestHandler {
	return &TestHandler{stateService}
}

func (listener *TestHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	text := message.Content.Text
	return types.MessageUnion{
		Text: strings.ToUpper(text),
	}
}
