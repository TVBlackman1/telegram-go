package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

type StateService struct {
	FirstState  *FirstState
	SecondState *SecondState
}

func NewStateService(rep *repository.Repository) *StateService {
	return &StateService{
		FirstState:  &FirstState{rep},
		SecondState: &SecondState{rep},
	}
}
