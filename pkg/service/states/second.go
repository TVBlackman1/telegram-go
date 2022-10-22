package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

const SECOND_STATE_NAME = "Second state"

type SecondState struct {
	commonContext *CommonStateContext
	queueMessages []types.MessageUnion
	notifications []notifier.NotifierContext
}

func NewSecondState(context *CommonStateContext) *SecondState {
	return &SecondState{
		commonContext: context,
		queueMessages: []types.MessageUnion{},
		notifications: []notifier.NotifierContext{},
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
		state.notifications = append(state.notifications, notifier.NotifierContext{
			ChatId: msg.Sender.ChatId,
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

func (state *SecondState) GetNotifications() []notifier.NotifierContext {
	return state.notifications
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
