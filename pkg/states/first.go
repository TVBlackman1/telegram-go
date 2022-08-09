package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FirstState struct {
	rep *repository.Repository
}

func (state *FirstState) PreparePresentation() types.MessageUnion {
	return types.MessageUnion{
		Text: "???",
	}
}

func (state *FirstState) ProcessUserInput(user *tgbotapi.User, msg types.MessageUnion) {
	if msg.Text == "2" {
		state.action(user, msg.Text)
	}
}

func (state *FirstState) action(user *tgbotapi.User, text string) {
	fmt.Printf("User %s sent: %s", user.FirstName, text)
}
