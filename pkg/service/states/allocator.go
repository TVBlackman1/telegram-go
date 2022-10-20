package states

import "github.com/TVBlackman1/telegram-go/pkg/repository"

func GetStateProcessor(stateName string, rep *repository.Repository) (retState UserState) {
	switch stateName {
	case FIRST_STATE_NAME:
		retState = NewFirstState(rep)
	case SECOND_STATE_NAME:
		retState = NewSecondState(rep)
	}
	return
}
