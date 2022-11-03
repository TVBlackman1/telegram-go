package second

import (
	"encoding/json"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
	thirdDto "github.com/TVBlackman1/telegram-go/pkg/service/state/states/third/dto"
)

func (state *SecondState) toThirdState(msg types.ReceivedMessage) {
	chatId := msg.Sender.ChatId
	stateSwitcher := state.commonContext.StateSwitcher
	newState := utils.StateOnFlyDto{
		Name:    utils.THIRD_STATE_NAME,
		Context: prepareThirdContext(),
	}
	stateSwitcher.TransferToNewStateByChatId(chatId, newState)
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: "Transfer to third state",
	})
}

func prepareThirdContext() string {
	dto := thirdDto.Context{
		JokesInRow: 0,
	}
	text, _ := json.Marshal(dto)
	return string(text)
}
