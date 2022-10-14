package states

import (
	"fmt"
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
)

const FIRST_STATE_NAME = "First state"

type FirstState struct {
	rep     *repository.Repository
	stateId uuid.UUID
}

func NewFirstState(rep *repository.Repository) *FirstState {
	return &FirstState{rep: rep}
}

func (state *FirstState) PreparePresentation() types.MessageUnion {
	return types.MessageUnion{
		Text: fmt.Sprintf("Your state is %s (%s)", FIRST_STATE_NAME, state.stateId),
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		fmt.Printf("User %s sent: %s", msg.Sender.Login, msg.Content.Text)
	}
}

func (state *FirstState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	stateId := fmt.Sprintf("%v", context)
	stateId = strings.ReplaceAll(stateId, "-", "")
	uuidId, err := uuid.Parse(stateId)
	if err != nil {
		return err
	}
	state.stateId = uuidId
	return nil
}

func (state *FirstState) GetName() string {
	return FIRST_STATE_NAME
}
