package impls

import (
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TestHandler struct {
	stateService *states.StateService
}

func NewTestHandler(stateService *states.StateService) *TestHandler {
	return &TestHandler{stateService}
}

func (listener *TestHandler) Process(message *tgbotapi.Message) types.MessageUnion {
	text := message.Text
	return types.MessageUnion{
		Text: strings.ToUpper(text),
	}

}
