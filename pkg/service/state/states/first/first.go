package first

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
)

type FirstState struct {
	commonContext *utils.CommonStateContext
	queueMessages []types.Message
	autoMessages  []notifier.NotifierContext
}

func NewFirstState(context *utils.CommonStateContext) *FirstState {
	return &FirstState{
		commonContext: context,
		queueMessages: []types.Message{},
		autoMessages:  []notifier.NotifierContext{},
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		state.toSecondState(msg)
	}
	state.autoMessages = lib.AddAutoMessageFromUserState(state.autoMessages, msg.Sender.ChatId)

}

func (state *FirstState) ProcessSystemInvoke(chatId types.ChatId) {
	state.queueMessages = lib.AddKeyboard(state.queueMessages, types.Keyboard{
		[]types.ButtonContent{"2"},
	})
}

func (state *FirstState) GetBotMessages() []types.Message {
	return state.queueMessages
}

func (state *FirstState) GetAutoMessages() []notifier.NotifierContext {
	return state.autoMessages
}

func (state *FirstState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
}

func (state *FirstState) SetState(msg types.ReceivedMessage, stateData repository.StateDbDto) error {
	return nil
}

func (state *FirstState) GetName() string {
	return utils.FIRST_STATE_NAME
}
