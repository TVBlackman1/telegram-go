package impls

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	stateService *states.StateService
}

func NewStartHandler(stateService *states.StateService) *StartHandler {
	return &StartHandler{stateService}
}

func (listener *StartHandler) Process(message *tgbotapi.Message) {
	fmt.Printf("Start from %s\n", message.From.UserName)

}
