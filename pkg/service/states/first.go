package states

import (
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/google/uuid"
)

const FIRST_STATE_NAME = "First state"

type FirstState struct {
	commonContext *CommonStateContext
	stateId       uuid.UUID
	toNewState    bool
	newStateInfo  string
}

func NewFirstState(context *CommonStateContext) *FirstState {
	return &FirstState{commonContext: context}
}

func (state *FirstState) PreparePresentation() types.MessageUnion {
	if state.toNewState {
		return types.MessageUnion{
			Text: state.newStateInfo,
		}
	}
	return types.MessageUnion{
		Text: fmt.Sprintf("Your state is %s (%s)", FIRST_STATE_NAME, state.stateId),
	}
}

func (state *FirstState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.toNewState = true
		state.newStateInfo = "Transfer to second state"
	}
}

func (state *FirstState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	stateId := fmt.Sprintf("%v", context)
	uuidId, err := uuid.Parse(stateId)
	if err != nil {
		return err
	}
	state.stateId = uuidId
	return nil
}

func (state *FirstState) GetName() string {
	return FIRST_STATE_NAME
}
