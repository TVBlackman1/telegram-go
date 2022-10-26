package lib

import (
	"github.com/TVBlackman1/telegram-go/pkg/constants"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

// TODO fix messages = ... to pointer logic
func AddKeyboard(messages []types.MessageUnion, keyboard types.Keyboard) {
	messages = append(messages, types.MessageUnion{
		Text:     constants.KEYBOARD_HAS_BEEN_OPENED,
		Keyboard: keyboard,
	})
}

func AddAutoMessageFromUserState(autoMessages []notifier.NotifierContext, chatId types.ChatId) {
	autoMessages = append(autoMessages, notifier.NotifierContext{
		ChatId: chatId,
	})
}
