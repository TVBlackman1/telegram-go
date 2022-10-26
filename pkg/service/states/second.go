package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/constants"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	commonContext *CommonStateContext
	queueMessages []types.MessageUnion
	autoMessages  []notifier.NotifierContext
}

func NewSecondState(context *CommonStateContext) *SecondState {
	return &SecondState{
		commonContext: context,
		queueMessages: []types.MessageUnion{},
		autoMessages:  []notifier.NotifierContext{},
	}
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
		state.autoMessages = append(state.autoMessages, notifier.NotifierContext{
			ChatId: msg.Sender.ChatId,
		})
	}
	if msg.Content.Text == "3" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{THIRD_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Transfer to third state",
		})
		state.autoMessages = append(state.autoMessages, notifier.NotifierContext{
			ChatId: msg.Sender.ChatId,
		})
	}
}

func (state *SecondState) ProcessSystemInvoke(chatId types.ChatId) {
	state.queueMessages = append(state.queueMessages, types.MessageUnion{
		Text: constants.KEYBOARD_HAS_BEEN_OPENED,
		Keyboard: types.Keyboard{
			[]types.ButtonContent{"1", "3"},
		},
	})
}

func (state *SecondState) GetBotMessages() []types.MessageUnion {
	return state.queueMessages
}

func (state *SecondState) GetAutoMessages() []notifier.NotifierContext {
	return state.autoMessages
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
