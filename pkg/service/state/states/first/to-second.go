package first

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
)

func (state *FirstState) toSecondState(msg types.ReceivedMessage) {
	chatId := msg.Sender.ChatId
	stateSwitcher := state.commonContext.StateSwitcher
	newState := utils.StateOnFlyDto{
		Name:    utils.SECOND_STATE_NAME,
		Context: "{}",
	}
	stateSwitcher.TransferToNewStateByChatId(chatId, newState)
	// notificator := state.commonContext.Notifier.GetNotificator()
	// defer func() { notificator <- notifier.NotifierContext{ChatId: chatId} }()
	// notificator <- notifier.NotifierContext{ChatId: chatId}
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: "Transfer to second state",
	})
}
