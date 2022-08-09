package presenter

import (
	"errors"
	"reflect"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Present(msg *tgbotapi.MessageConfig, data types.MessageUnion) error {
	if data.Text != "" {
		msg.Text = data.Text
	}
	if !reflect.ValueOf(data.Keyboard).IsZero() {
		replyKeyboard := tgbotapi.NewReplyKeyboard()
		for i := 0; i < len(data.Keyboard); i++ {
			replyKeyboard.Keyboard = append(replyKeyboard.Keyboard, tgbotapi.NewKeyboardButtonRow())
			for j := 0; j < len(data.Keyboard[i]); j++ {
				text := string(data.Keyboard[i][j])
				lastRow := replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1]
				lastRow = append(lastRow, tgbotapi.NewKeyboardButton(text))
				replyKeyboard.Keyboard[len(replyKeyboard.Keyboard)-1] = lastRow
			}
		}
	}
	if len(data.Media) > 0 {
		return errors.New("media is not supported")
	}
	return nil
}

func Collect(msg *tgbotapi.MessageConfig) types.MessageUnion {
	var messageDto types.MessageUnion
	if msg.Text != "" {
		messageDto.Text = msg.Text
	}
	return messageDto
}
