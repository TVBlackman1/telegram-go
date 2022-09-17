package service

import (
	"errors"
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
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
		ChatId: message.Sender.ChatId,
	})
}

func (stateService *StateService) RegisterNewUser(sender types.Sender) types.MessageUnion {
	var retMessage types.MessageUnion
	newUser := repository.CreateUserDto{
		Id:     uuid.New(),
		Login:  sender.Login,
		Name:   sender.Name,
		ChatId: sender.ChatId,
	}
	_, err := stateService.rep.UserRepository.Add(newUser)
	if err != nil && errors.Is(err, utils.ErrAlreadyExists) {
		retMessage = types.MessageUnion{
			Text: fmt.Sprintf("You are already using this bot, %s!", sender.Name),
		}
	} else {
		retMessage = types.MessageUnion{
			Text: fmt.Sprintf("Thanks for using bot, %s!", sender.Name),
		}
	}
	return retMessage
}
