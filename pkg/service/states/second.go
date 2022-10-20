package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	rep          *repository.Repository
	toNewState   bool
	newStateInfo string
}

func NewSecondState(rep *repository.Repository) *SecondState {
	return &SecondState{rep: rep}
}

func (state *SecondState) PreparePresentation() types.MessageUnion {
	if state.toNewState {
		return types.MessageUnion{
			Text: state.newStateInfo,
		}
	}
	return types.MessageUnion{
		Text: "this is second state",
	}
}

func (state *SecondState) ProcessUserInput(msg types.ReceivedMessage) {
	// TODO refactor below
	if msg.Content.Text == "1" {
		// support errors
		user, _ := state.rep.UserRepository.GetOne(repository.UserQuery{
			ChatId: msg.Sender.ChatId,
		})
		newStateId, _ := state.rep.StateRepository.Add(repository.CreateStateDto{
			Id:      uuid.New(),
			Name:    FIRST_STATE_NAME,
			Context: "{}",
		})
		state.rep.UserRepository.SetNewStateUUID(user.Id, newStateId)
		state.toNewState = true
		state.newStateInfo = "Transfer to first state"
	}
	if msg.Content.Text == "3" {
		// support errors
		user, _ := state.rep.UserRepository.GetOne(repository.UserQuery{
			ChatId: msg.Sender.ChatId,
		})
		newStateId, _ := state.rep.StateRepository.Add(repository.CreateStateDto{
			Id:      uuid.New(),
			Name:    THIRD_STATE_NAME,
			Context: "{}",
		})
		state.rep.UserRepository.SetNewStateUUID(user.Id, newStateId)
		state.toNewState = true
		state.newStateInfo = "Transfer to third state"
	}
}

func (state *SecondState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	return nil
}

func (state *SecondState) action(login string, text string) {
	fmt.Printf("User %s sent: %s", login, text)
}

func (state *SecondState) GetName() string {
	return SECOND_STATE_NAME
}
