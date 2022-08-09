package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserState interface {
	PreparePresentation() types.MessageUnion
	ProcessUserInput(user *tgbotapi.User, msg types.MessageUnion)
}
