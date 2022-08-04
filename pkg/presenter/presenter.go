package presenter

import (
	"errors"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Present(msg *tgbotapi.MessageConfig, data MessageUnion) error {
	if data.text != "" {
		msg.Text = data.text
	}
	if !reflect.ValueOf(data.keyboard).IsZero() {
		replyKeyboard := tgbotapi.NewReplyKeyboard()
		for i := 0; i < len(data.keyboard); i++ {
			replyKeyboard.Keyboard = append(replyKeyboard.Keyboard, tgbotapi.NewKeyboardButtonRow())
			for j := 0; j < len(data.keyboard[i]); j++ {
				text := string(data.keyboard[i][j])
				lastRow := replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1]
				lastRow = append(lastRow, tgbotapi.NewKeyboardButton(text))
				replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1] = lastRow
			}
		}
	}
	if len(data.media) > 0 {
		return errors.New("Media is not supported")
	}
	return nil
}
