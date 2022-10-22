package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	commonContext *CommonStateContext
	queueMessages []types.MessageUnion
}

func NewSecondState(context *CommonStateContext) *SecondState {
	return &SecondState{commonContext: context}
}

func (state *SecondState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "1" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{FIRST_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Transfer to first state",
		})
	}
	if msg.Content.Text == "3" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{THIRD_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		fmt.Println("Before adding")
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Transfer to third state",
		})
	}
}

func (state *SecondState) ProcessSystemInvoke(chatId types.ChatId) {
	panic("not implemented")
}

func (state *SecondState) GetBotMessages() []types.MessageUnion {
	return state.queueMessages
}

func (state *SecondState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
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
