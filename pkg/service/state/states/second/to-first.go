package second

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
)

func (state *SecondState) toFirstState(msg types.ReceivedMessage) {
	chatId := msg.Sender.ChatId
	stateSwitcher := state.commonContext.StateSwitcher
	newState := utils.StateOnFlyDto{
		Name:    utils.FIRST_STATE_NAME,
		Context: "{}",
	}
	stateSwitcher.TransferToNewStateByChatId(chatId, newState)
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: "Transfer to first state",
	})
}
