package states

import (
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

type CommonStateContext struct {
	rep           *repository.Repository
	StateSwitcher *StateSwitcher
	Notifier      *notifier.SystemNotifier
}

func NewCommonStateContext(rep *repository.Repository, notifier *notifier.SystemNotifier) *CommonStateContext {
	return &CommonStateContext{
		rep:           rep,
		StateSwitcher: NewStateSwitcher(rep),
		Notifier:      notifier,
	}
}
