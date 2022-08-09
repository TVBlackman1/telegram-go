package impls

import (
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/states"
)

type TestHandler struct {
	stateService *states.StateService
}

func NewTestHandler(stateService *states.StateService) *TestHandler {
	return &TestHandler{stateService}
}

func (listener *TestHandler) Process(message types.ReceivedMessage) types.MessageUnion {
	text := message.Content.Text
	return types.MessageUnion{
		Text: strings.ToUpper(text),
	}

}
