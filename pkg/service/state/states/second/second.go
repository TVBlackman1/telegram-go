package second

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
)

type SecondState struct {
	commonContext *utils.CommonStateContext
	queueMessages []types.Message
	autoMessages  []notifier.NotifierContext
}

func NewSecondState(context *utils.CommonStateContext) *SecondState {
	return &SecondState{
		commonContext: context,
		queueMessages: []types.Message{},
		autoMessages:  []notifier.NotifierContext{},
	}
}

func (state *SecondState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "1" {
		state.toFirstState(msg)
	}
	if msg.Content.Text == "3" {
		state.toThirdState(msg)
	}
	state.autoMessages = lib.AddAutoMessageFromUserState(state.autoMessages, msg.Sender.ChatId)

}

func (state *SecondState) ProcessSystemInvoke(chatId types.ChatId) {
	state.queueMessages = lib.AddKeyboard(state.queueMessages, types.Keyboard{
		[]types.ButtonContent{"1", "3"},
	})
}

func (state *SecondState) GetBotMessages() []types.Message {
	return state.queueMessages
}

func (state *SecondState) GetAutoMessages() []notifier.NotifierContext {
	return state.autoMessages
}

func (state *SecondState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
}

func (state *SecondState) SetState(msg types.ReceivedMessage, stateData repository.StateDbDto) error {
	return nil
}

func (state *SecondState) GetName() string {
	return utils.SECOND_STATE_NAME
}
