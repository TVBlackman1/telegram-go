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
	// TODO rename stateService to common variant
	stateService.rep.UserRepository.GetOne(repository.UserQuery{
		ChatId: message.Sender.ChatId,
	})
}

func (stateService *StateService) RegisterNewUser(sender types.Sender) (retMessage types.MessageUnion) {
	newUser := repository.CreateUserDto{
		Id:     uuid.New(),
		Login:  sender.Login,
		Name:   sender.Name,
		ChatId: sender.ChatId,
	}
	userUUID, err := stateService.rep.UserRepository.Add(newUser)
	if err != nil {
		var textForSending string
		// TODO change error processing
		if errors.Is(err, utils.ErrAlreadyExists) {
			textForSending = fmt.Sprintf("You are already using this bot, %s!", sender.Name)
		} else {
			textForSending = "Some error. Try again later"
		}
		retMessage = types.MessageUnion{
			Text: textForSending,
		}
		return
	}
	stateUUID, err := stateService.createNewDefaultState()
	if err != nil {
		retMessage = types.MessageUnion{
			Text: "Some error. Try again later",
		}
		return
	}
	err = stateService.rep.UserRepository.SetNewStateUUID(userUUID, stateUUID)
	if err != nil {
		retMessage = types.MessageUnion{
			Text: "Some error. Try again later",
		}
		return
	}
	retMessage = types.MessageUnion{
		Text: fmt.Sprintf("Thanks for using bot, %s!", sender.Name),
	}
	return
}

func (stateService *StateService) createNewDefaultState() (uuid.UUID, error) {
	// TODO change context to normal usability
	defaultState := repository.CreateStateDto{
		Id:      uuid.New(),
		Name:    states.FIRST_STATE_NAME,
		Context: "{}",
	}

	// TODO check why cant use := and must =
	return stateService.rep.StateRepository.Add(defaultState)
}
