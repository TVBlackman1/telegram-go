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

type UserService struct {
	rep          *repository.Repository
	stateContext *states.CommonStateContext
}

func NewUserService(rep *repository.Repository) *UserService {
	stateContext := states.NewCommonStateContext(rep)
	return &UserService{rep, stateContext}
}

func (userService *UserService) GetCurrentState(chatId types.ChatId) (repository.StateDbDto, error) {
	user, err := userService.rep.UserRepository.GetOne(repository.UserQuery{
		ChatId: chatId,
	})
	if err != nil {
		return repository.StateDbDto{}, err
	}
	return user.State, nil
}

func (userService *UserService) GetCurrentStateProcessor(currentState repository.StateDbDto) (states.UserState, error) {
	stateProcessor := states.GetStateProcessor(currentState.Name, userService.stateContext)
	return stateProcessor, nil
}

func (userService *UserService) RegisterNewUser(sender types.Sender) (retMessage types.MessageUnion) {
	newUser := repository.CreateUserDto{
		Id:     uuid.New(),
		Login:  sender.Login,
		Name:   sender.Name,
		ChatId: sender.ChatId,
	}
	userUUID, err := userService.rep.UserRepository.Add(newUser)
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
	stateUUID, err := userService.createNewDefaultState()
	if err != nil {
		retMessage = types.MessageUnion{
			Text: "Some error. Try again later",
		}
		return
	}
	err = userService.rep.UserRepository.SetNewStateUUID(userUUID, stateUUID)
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

func (userService *UserService) createNewDefaultState() (uuid.UUID, error) {
	// TODO change context to normal usability
	defaultState := repository.CreateStateDto{
		Id:      uuid.New(),
		Name:    states.FIRST_STATE_NAME,
		Context: "{}",
	}

	// TODO check why cant use := and must =
	return userService.rep.StateRepository.Add(defaultState)
}
