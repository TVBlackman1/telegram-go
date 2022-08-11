package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

type FirstState struct {
	rep *repository.Repository
}

func NewFirstState(rep *repository.Repository) *FirstState {
	return &FirstState{rep}
}

func (state *FirstState) PreparePresentation() types.MessageUnion {
	return types.MessageUnion{
		Text: "???",
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		fmt.Printf("User %s sent: %s", msg.Sender.Login, msg.Content.Text)
	}
}

func (state *FirstState) SetState(sender types.Sender, context interface{}) {
	// TODO
}
