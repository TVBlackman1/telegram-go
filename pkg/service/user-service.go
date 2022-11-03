package service

import (
	"errors"
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
	"github.com/TVBlackman1/telegram-go/pkg/service/state"
	stateUtils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states"
	"github.com/google/uuid"
)

type UserService struct {
	Rep          *repository.Repository
	StateContext *stateUtils.CommonStateContext
}

func NewUserService(rep *repository.Repository, notifier *notifier.SystemNotifier) *UserService {
	stateContext := stateUtils.NewCommonStateContext(rep, notifier)
	return &UserService{rep, stateContext}
}

func (userService *UserService) GetCurrentState(chatId types.ChatId) (repository.StateDbDto, error) {
	user, err := userService.Rep.UserRepository.GetOne(repository.UserQuery{
		ChatId: chatId,
	})
	if err != nil {
		return repository.StateDbDto{}, err
	}
	return user.State, nil
}

func (userService *UserService) GetCurrentStateProcessor(currentState repository.StateDbDto) (states.UserState, error) {
	stateProcessor := state.GetStateProcessor(currentState.Name, userService.StateContext)
	return stateProcessor, nil
}

func (userService *UserService) RegisterNewUser(sender types.Sender) (retMessage types.Message) {
	newUser := repository.CreateUserDto{
		Id:     uuid.New(),
		Login:  sender.Login,
		Name:   sender.Name,
		ChatId: sender.ChatId,
	}
	userUUID, err := userService.Rep.UserRepository.Add(newUser)
	if err != nil {
		var textForSending string
		// TODO change error processing
		if errors.Is(err, utils.ErrAlreadyExists) {
			textForSending = fmt.Sprintf("You are already using this bot, %s!", sender.Name)
		} else {
			textForSending = "Some error. Try again later"
		}
		retMessage = types.Message{
			Text: textForSending,
		}
		return
	}
	defaultState := stateUtils.StateOnFlyDto{Name: stateUtils.FIRST_STATE_NAME, Context: "{}"}
	userService.StateContext.StateSwitcher.TransferToNewState(userUUID, defaultState)
	retMessage = types.Message{
		Text: fmt.Sprintf("Thanks for using bot, %s!", sender.Name),
	}
	return
}
