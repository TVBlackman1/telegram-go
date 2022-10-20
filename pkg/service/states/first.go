package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
)

const FIRST_STATE_NAME = "First state"

type FirstState struct {
	rep          *repository.Repository
	stateId      uuid.UUID
	toNewState   bool
	newStateInfo string
}

func NewFirstState(rep *repository.Repository) *FirstState {
	return &FirstState{rep: rep}
}

func (state *FirstState) PreparePresentation() types.MessageUnion {
	if state.toNewState {
		return types.MessageUnion{
			Text: state.newStateInfo,
		}
	}
	return types.MessageUnion{
		Text: fmt.Sprintf("Your state is %s (%s)", FIRST_STATE_NAME, state.stateId),
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		// support errors
		user, _ := state.rep.UserRepository.GetOne(repository.UserQuery{
			ChatId: msg.Sender.ChatId,
		})
		newStateId, _ := state.rep.StateRepository.Add(repository.CreateStateDto{
			Id:      uuid.New(),
			Name:    SECOND_STATE_NAME,
			Context: "{}",
		})
		state.rep.UserRepository.SetNewStateUUID(user.Id, newStateId)
		state.toNewState = true
		state.newStateInfo = "Transfer to second state"
	}
}

func (state *FirstState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	stateId := fmt.Sprintf("%v", context)
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
