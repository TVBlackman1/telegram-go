package impls

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/service"
)

type StartHandler struct {
	stateService *service.StateService
}

func NewStartHandler(stateService *service.StateService) *StartHandler {
	return &StartHandler{stateService}
}

func (listener *StartHandler) Process(message types.ReceivedMessage) {
	fmt.Printf("Start from %d\n", message.ChatId)

}
