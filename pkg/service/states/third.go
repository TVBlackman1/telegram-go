package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

const THIRD_STATE_NAME = "Third state"

type ThirdState struct {
	commonContext *CommonStateContext
	jokeToAnswer  bool
	toNewState    bool
	newStateInfo  string
}

func NewThirdState(context *CommonStateContext) *ThirdState {
	return &ThirdState{commonContext: context}
}

func (state *ThirdState) PreparePresentation() types.MessageUnion {
	if state.toNewState {
		return types.MessageUnion{
			Text: state.newStateInfo,
		}
	}
	return types.MessageUnion{
		Text: "this is third state",
	}
}

func (state *ThirdState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		// support errors
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.toNewState = true
		state.newStateInfo = "Transfer to second state"

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
