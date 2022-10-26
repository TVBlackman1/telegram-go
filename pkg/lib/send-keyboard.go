package lib

import (
	"github.com/TVBlackman1/telegram-go/pkg/constants"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
)

// TODO maybe later
func SendKeyboard(messages []types.MessageUnion, keyboard types.Keyboard) {
	messages = append(messages, types.MessageUnion{
		Text:     constants.KEYBOARD_HAS_BEEN_OPENED,
		Keyboard: keyboard,
	})
}
