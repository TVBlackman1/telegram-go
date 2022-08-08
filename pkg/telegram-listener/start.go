package telegramlistener

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartListener struct {
	stateService *states.StateService
}

func (listener *StartListener) Process(message *tgbotapi.Message) {
	fmt.Printf("Start from %s\n", message.From.UserName)

}
