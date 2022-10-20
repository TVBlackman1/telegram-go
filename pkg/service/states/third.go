package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
)

const THIRD_STATE_NAME = "Third state"

type ThirdState struct {
	rep          *repository.Repository
	jokeToAnswer bool
}

func NewThirdState(rep *repository.Repository) *ThirdState {
	return &ThirdState{
		rep:          rep,
		jokeToAnswer: false,
	}
}

func (state *ThirdState) PreparePresentation() types.MessageUnion {
	if state.jokeToAnswer {
		return types.MessageUnion{
			Text: "Смешная шутка не готова",
		}
	}
	return types.MessageUnion{
		Text: "this is third state",
	}
}

func (state *ThirdState) ProcessUserInput(msg types.ReceivedMessage) {
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
		fmt.Println("User changed state to 2")
	} else {
		state.jokeToAnswer = true
	}
}

func (state *ThirdState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	return nil
}

func (state *ThirdState) action(login string, text string) {
	fmt.Printf("User %s sent: %s", login, text)
}

func (state *ThirdState) GetName() string {
	return THIRD_STATE_NAME
}
