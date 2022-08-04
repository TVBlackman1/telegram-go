package states

import "github.com/TVBlackman1/telegram-go/pkg/presenter"

type FirstState struct{}

func (state *FirstState) PreparePresentation() presenter.MessageUnion {
	return presenter.MessageUnion{}
}

func (state *FirstState) ProcessUserInput(msg presenter.MessageUnion) {

}
