package service

import (
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/service/states"
)

type StateService struct {
	FirstState  *states.FirstState
	SecondState *states.SecondState
}

func NewStateService(rep *repository.Repository) *StateService {
	return &StateService{
		FirstState:  states.NewFirstState(rep),
		SecondState: states.NewSecondState(rep),
	}
}
