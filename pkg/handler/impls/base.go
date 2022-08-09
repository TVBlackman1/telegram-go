package impls

import (
	"github.com/TVBlackman1/telegram-go/pkg/presenter/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ConcreteHandler interface {
	Process(message *tgbotapi.Message) types.MessageUnion
}
