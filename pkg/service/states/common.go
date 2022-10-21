package states

import "github.com/TVBlackman1/telegram-go/pkg/repository"

type CommonStateContext struct {
	rep           *repository.Repository
	StateSwitcher *StateSwitcher
}

func NewCommonStateContext(rep *repository.Repository) *CommonStateContext {
	return &CommonStateContext{
		rep:           rep,
		StateSwitcher: NewStateSwitcher(rep),
	}
}
