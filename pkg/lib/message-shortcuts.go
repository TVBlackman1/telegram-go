package lib

import (
	"github.com/TVBlackman1/telegram-go/pkg/constants"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
)

func AddKeyboard(messages []types.Message, keyboard types.Keyboard) []types.Message {
	return append(messages, types.Message{
		Text:     constants.KEYBOARD_HAS_BEEN_OPENED,
		Keyboard: keyboard,
	})
}

func AddAutoMessageFromUserState(autoMessages []notifier.NotifierContext, chatId types.ChatId) []notifier.NotifierContext {
	return append(autoMessages, notifier.NotifierContext{
		ChatId: chatId,
	})
}
