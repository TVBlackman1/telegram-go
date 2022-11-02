package presenter

import (
	"errors"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Present(msg *tgbotapi.MessageConfig, data types.Message) error {
	if data.Text != "" {
		msg.Text = data.Text
	}
	if len(data.Keyboard) > 0 && len(data.Text) > 0 {
		replyKeyboard := convertKeyboard(data.Keyboard)
		msg.ReplyMarkup = replyKeyboard
	} else {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}
	if len(data.Media) > 0 {
		return errors.New("media is not supported")
	}
	return nil
}

func Collect(msg *tgbotapi.Message) types.Message {
	var messageDto types.Message
	if msg.Text != "" {
		messageDto.Text = msg.Text
	}
	return messageDto
}

func convertKeyboard(keyboard types.Keyboard) tgbotapi.ReplyKeyboardMarkup {
	replyKeyboard := tgbotapi.NewReplyKeyboard()
	for i := 0; i < len(keyboard); i++ {
		replyKeyboard.Keyboard = append(replyKeyboard.Keyboard, tgbotapi.NewKeyboardButtonRow())
		for j := 0; j < len(keyboard[i]); j++ {
			text := string(keyboard[i][j])
			lastRow := replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1]
			lastRow = append(lastRow, tgbotapi.NewKeyboardButton(text))
			replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1] = lastRow
		}
	}
	return replyKeyboard
}
