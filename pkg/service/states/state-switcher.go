package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
)

type StateOnFlyDto struct {
	Name    string
	Context string
}

// func SetStateByUserId(userId uuid.UUID) {

// }

// func SetStateByChatId(chatId types.ChatId) {
// 	user, _ := state.rep.UserRepository.GetOne(repository.UserQuery{
// 		ChatId: msg.Sender.ChatId,
// 	})
// }

// func TransferToNewState(state UserState) {

// }

type StateSwitcher struct {
	rep *repository.Repository
}

func NewStateSwitcher(rep *repository.Repository) *StateSwitcher {
	return &StateSwitcher{rep}
}

func (stateSwitcher *StateSwitcher) TransferToNewState(userId uuid.UUID, state StateOnFlyDto) error {
	// support errors
	newStateId, _ := stateSwitcher.rep.StateRepository.Add(repository.CreateStateDto{
		Id:      uuid.New(),
		Name:    state.Name,
		Context: state.Context,
	})
	return stateSwitcher.rep.UserRepository.SetNewStateUUID(userId, newStateId)
}

func (stateSwitcher *StateSwitcher) TransferToNewStateByChatId(chatId types.ChatId, state StateOnFlyDto) error {
	// support errors
	user, _ := stateSwitcher.rep.UserRepository.GetOne(repository.UserQuery{
		ChatId: chatId,
	})
	return stateSwitcher.TransferToNewState(user.Id, state)
}
