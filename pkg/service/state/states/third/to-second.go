package third

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
)

func (state *ThirdState) toSecondState(msg types.ReceivedMessage) {
	// support errors
	chatId := msg.Sender.ChatId
	stateSwitcher := state.commonContext.StateSwitcher
	newState := utils.StateOnFlyDto{
		Name:    utils.SECOND_STATE_NAME,
		Context: "{}",
	}
	stateSwitcher.TransferToNewStateByChatId(chatId, newState)
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: "Transfer to second state",
	})
}
