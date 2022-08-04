package states

import "github.com/TVBlackman1/telegram-go/pkg/presenter"

type UserState interface {
	PreparePresentation() presenter.MessageUnion
	ProcessUserInput(msg presenter.MessageUnion)
}
