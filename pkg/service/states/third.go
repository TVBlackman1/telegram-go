package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

const THIRD_STATE_NAME = "Third state"

type ThirdState struct {
	commonContext *CommonStateContext
	jokeToAnswer  bool
	queueMessages []types.MessageUnion
	autoMessages  []notifier.NotifierContext
}

func NewThirdState(context *CommonStateContext) *ThirdState {
	return &ThirdState{
		commonContext: context,
		queueMessages: []types.MessageUnion{},
		autoMessages:  []notifier.NotifierContext{},
	}
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
	if msg.Content.Text == "joke" {
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Sorry. I do not know jokes",
		})
	}
	lib.AddAutoMessageFromUserState(state.autoMessages, msg.Sender.ChatId)

}

func (state *ThirdState) ProcessSystemInvoke(chatId types.ChatId) {
	lib.AddKeyboard(state.queueMessages, types.Keyboard{
		[]types.ButtonContent{"2", "joke"},
	})
}

func (state *ThirdState) GetBotMessages() []types.MessageUnion {
	return state.queueMessages
}

func (state *ThirdState) GetAutoMessages() []notifier.NotifierContext {
	return state.autoMessages
}

func (state *ThirdState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
}

func (state *ThirdState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	return nil
}

func (state *ThirdState) GetName() string {
	return THIRD_STATE_NAME
}
