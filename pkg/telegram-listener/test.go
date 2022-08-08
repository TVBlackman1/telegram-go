package telegramlistener

import (
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	"github.com/TVBlackman1/telegram-go/pkg/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TestListener struct {
	stateService *states.StateService
}

func (listener *TestListener) Process(message *tgbotapi.Message) presenter.MessageUnion {
	text := message.Text
	return presenter.MessageUnion{
		Text: strings.ToUpper(text),
	}

}
