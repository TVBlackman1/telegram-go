package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FirstState struct{}

func (state *FirstState) PreparePresentation() presenter.MessageUnion {
	return presenter.MessageUnion{
		Text: "???",
	}
}

func (state *FirstState) ProcessUserInput(user *tgbotapi.User, msg presenter.MessageUnion) {
	if msg.Text == "2" {
		state.action(user, msg.Text)
	}
}

func (state *FirstState) action(user *tgbotapi.User, text string) {
	fmt.Printf("User %s sent: %s", user.FirstName, text)
}
