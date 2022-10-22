package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/google/uuid"
)

const FIRST_STATE_NAME = "First state"

type FirstState struct {
	commonContext *CommonStateContext
	stateId       uuid.UUID
	queueMessages []types.MessageUnion // TODO change to chan notifier.NotifierContext
}

func NewFirstState(context *CommonStateContext) *FirstState {
	return &FirstState{
		commonContext: context,
		queueMessages: []types.MessageUnion{},
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	state.queueMessages = append(state.queueMessages, types.MessageUnion{
		Text: "Current state is first",
	})
	if msg.Content.Text == "2" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		// notificator := state.commonContext.Notifier.GetNotificator()
		// defer func() { notificator <- notifier.NotifierContext{ChatId: chatId} }()
		// notificator <- notifier.NotifierContext{ChatId: chatId}
		state.queueMessages = append(state.queueMessages, types.MessageUnion{
			Text: "Transfer to second state",
		})
	}
}

func (state *FirstState) ProcessSystemInvoke(chatId types.ChatId) {
	panic("not implemented")
}

func (state *FirstState) GetBotMessages() []types.MessageUnion {
	return state.queueMessages
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
