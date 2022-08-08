package telegramlistener

import "github.com/TVBlackman1/telegram-go/pkg/states"

type Handler struct {
	StartListener  *StartListener
	StatesListener *StatesListener
	TestListener   *TestListener
}

func NewHandler(stateService *states.StateService) *Handler {
	return &Handler{
		StartListener:  &StartListener{stateService},
		StatesListener: &StatesListener{stateService},
		TestListener:   &TestListener{stateService},
	}
}
