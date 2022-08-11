package handlers

import (
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type TestHandler struct {
	stateService *service.StateService
}

func NewTestHandler(stateService *service.StateService) *TestHandler {
	return &TestHandler{stateService}
}

func (listener *TestHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	text := message.Content.Text
	return types.MessageUnion{
		Text: strings.ToUpper(text),
	}
}
