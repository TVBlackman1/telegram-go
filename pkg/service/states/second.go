package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	commonContext *CommonStateContext
	toNewState    bool
	newStateInfo  string
}

func NewSecondState(context *CommonStateContext) *SecondState {
	return &SecondState{commonContext: context}
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
	if msg.Content.Text == "1" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{FIRST_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.toNewState = true
		state.newStateInfo = "Transfer to first state"
	}
	if msg.Content.Text == "3" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{THIRD_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
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
