package impls

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/states"
)

type StartHandler struct {
	stateService *states.StateService
}

func NewStartHandler(stateService *states.StateService) *StartHandler {
	return &StartHandler{stateService}
}

func (listener *StartHandler) Process(message types.ReceivedMessage) {
	fmt.Printf("Start from %d\n", message.ChatId)

}
