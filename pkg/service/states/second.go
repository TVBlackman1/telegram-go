package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	rep *repository.Repository
}

func NewSecondState(rep *repository.Repository) *SecondState {
	return &SecondState{rep}
}

func (state *SecondState) PreparePresentation() types.MessageUnion {
	return types.MessageUnion{
		Text: "???",
	}
}

func (state *SecondState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		state.action(msg.Sender.Login, msg.Content.Text)
	}
}

func (state *SecondState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	panic("not implemented")
}

func (state *SecondState) action(login string, text string) {
	fmt.Printf("User %s sent: %s", login, text)
}

func (state *SecondState) GetName() string {
	return SECOND_STATE_NAME
}
