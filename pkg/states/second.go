package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SecondState struct {
	rep *repository.Repository
}

func (state *SecondState) PreparePresentation() presenter.MessageUnion {
	return presenter.MessageUnion{
		Text: "???",
	}
}

func (state *SecondState) ProcessUserInput(user *tgbotapi.User, msg presenter.MessageUnion) {
	if msg.Text == "2" {
		state.action(user, msg.Text)
	}
}

func (state *SecondState) action(user *tgbotapi.User, text string) {
	fmt.Printf("User %s sent: %s", user.FirstName, text)
}
