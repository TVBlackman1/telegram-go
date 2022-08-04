package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserState interface {
	PreparePresentation() presenter.MessageUnion
	ProcessUserInput(user *tgbotapi.User, msg presenter.MessageUnion)
}
