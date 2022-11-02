package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/google/uuid"
)

const FIRST_STATE_NAME = "First state"

type FirstState struct {
	commonContext *CommonStateContext
	stateId       uuid.UUID
	queueMessages []types.Message
	autoMessages  []notifier.NotifierContext
}

func NewFirstState(context *CommonStateContext) *FirstState {
	return &FirstState{
		commonContext: context,
		queueMessages: []types.Message{},
		autoMessages:  []notifier.NotifierContext{},
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		// notificator := state.commonContext.Notifier.GetNotificator()
		// defer func() { notificator <- notifier.NotifierContext{ChatId: chatId} }()
		// notificator <- notifier.NotifierContext{ChatId: chatId}
		state.queueMessages = append(state.queueMessages, types.Message{
			Text: "Transfer to second state",
		})
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

func (state *FirstState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	stateId := fmt.Sprintf("%v", context)
	uuidId, err := uuid.Parse(stateId)
	if err != nil {
		return err
	}
	state.stateId = uuidId
	return nil
}

func (state *FirstState) GetName() string {
	return FIRST_STATE_NAME
}
