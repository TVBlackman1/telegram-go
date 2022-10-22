package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

const THIRD_STATE_NAME = "Third state"

type ThirdState struct {
	commonContext *CommonStateContext
	jokeToAnswer  bool
	queueMessages []types.MessageUnion // TODO change to chan notifier.NotifierContext
}

func NewThirdState(context *CommonStateContext) *ThirdState {
	return &ThirdState{commonContext: context}
}

func (state *ThirdState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		// support errors
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Transfer to second state",
		})
	}
}

func (state *ThirdState) ProcessSystemInvoke(chatId types.ChatId) {
	panic("not implemented")
}

func (state *ThirdState) GetBotMessages() []types.MessageUnion {
	return state.queueMessages
}

func (state *ThirdState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
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
