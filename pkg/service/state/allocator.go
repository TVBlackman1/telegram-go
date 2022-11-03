package state

import (
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states/first"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states/second"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states/third"
)

func GetStateProcessor(stateName string, stateContext *utils.CommonStateContext) (retState states.UserState) {
	switch stateName {
	case utils.FIRST_STATE_NAME:
		retState = first.NewFirstState(stateContext)
	case utils.SECOND_STATE_NAME:
		retState = second.NewSecondState(stateContext)
	case utils.THIRD_STATE_NAME:
		retState = third.NewThirdState(stateContext)
	}
	return
}
