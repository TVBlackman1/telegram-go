package service

import (
	"fmt"
	"os"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/service/states"
	"github.com/google/uuid"
)

type StateService struct {
	FirstState  *states.FirstState
	SecondState *states.SecondState
	rep         *repository.Repository
}

func NewStateService(rep *repository.Repository) *StateService {
	return &StateService{
		FirstState:  states.NewFirstState(rep),
		SecondState: states.NewSecondState(rep),
		rep:         rep,
	}
}

func (stateService *StateService) GetCurrentState(message types.ReceivedMessage) {
	stateService.rep.GetOne(repository.UserQuery{
		ChatId: message.ChatId,
	})
}

func (stateService *StateService) RegisterNewUser(chatId types.ChatId) {
	newUser := repository.CreateUserDto{
		Id:     uuid.New(),
		Name:   "",
		ChatId: chatId,
	}
	_, err := stateService.rep.UserRepository.Add(newUser)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
