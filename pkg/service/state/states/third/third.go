package third

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	utils "github.com/TVBlackman1/telegram-go/pkg/service/state/state-inner-utils"
	"github.com/TVBlackman1/telegram-go/pkg/service/state/states/third/dto"
	"github.com/google/uuid"
)

// TODO transfer all big methods to other file.go in this package
type ThirdState struct {
	commonContext *utils.CommonStateContext
	queueMessages []types.Message
	autoMessages  []notifier.NotifierContext
	context       dto.Context
	stateId       uuid.UUID
}

func NewThirdState(context *utils.CommonStateContext) *ThirdState {
	return &ThirdState{
		commonContext: context,
		queueMessages: []types.Message{},
		autoMessages:  []notifier.NotifierContext{},
	}
}

func (state *ThirdState) ProcessUserInput(msg types.ReceivedMessage) {
	if msg.Content.Text == "2" {
		// support errors
		chatId := msg.Sender.ChatId
		stateSwitcher := state.commonContext.StateSwitcher
		newState := utils.StateOnFlyDto{
			Name:    utils.SECOND_STATE_NAME,
			Context: "{}",
		}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.queueMessages = append(state.queueMessages, types.Message{
			Text: "Transfer to second state",
		})
	}
	if msg.Content.Text == "joke" {
		// TODO refactor
		if state.context.JokesInRow > 4 {
			state.queueMessages = append(state.queueMessages, types.Message{
				Text: "Stop laught! Take a break",
			})
			state.autoMessages = lib.AddAutoMessageFromUserState(state.autoMessages, msg.Sender.ChatId)
			return
		}
		jokeCount, _ := state.commonContext.Rep.JokeRepository.Count("")
		if jokeCount == 0 {
			state.queueMessages = append(state.queueMessages, types.Message{
				Text: "Sorry. I do not know interesting jokes...",
			})
		} else {
			randomJokeNumber := rand.Int() % jokeCount
			joke, _ := state.commonContext.Rep.JokeRepository.GetOne(repository.JokeQuery{
				Offset: randomJokeNumber,
			})
			state.queueMessages = append(state.queueMessages, types.Message{
				Text: joke.Text,
			})
		}
		state.context.JokesInRow++
		newContext, _ := json.Marshal(state.context)
		state.commonContext.Rep.JokeRepository.UpdateContext(state.stateId, string(newContext))
	}
	state.autoMessages = lib.AddAutoMessageFromUserState(state.autoMessages, msg.Sender.ChatId)

}

func (state *ThirdState) ProcessSystemInvoke(chatId types.ChatId) {
	state.queueMessages = lib.AddKeyboard(state.queueMessages, types.Keyboard{
		[]types.ButtonContent{"2", "joke"},
	})
}

func (state *ThirdState) GetBotMessages() []types.Message {
	return state.queueMessages
}

func (state *ThirdState) GetAutoMessages() []notifier.NotifierContext {
	return state.autoMessages
}

func (state *ThirdState) ProcessContextedSystemInvoke(chatId types.ChatId, context interface{}) {
	panic("not implemented")
}

func (state *ThirdState) SetState(msg types.ReceivedMessage, stateData repository.StateDbDto) error {
	// TODO add error checker
	json.Unmarshal([]byte(stateData.Context), &state.context)
	fmt.Printf("state.context.JokesInRow: %v\n", state.context.JokesInRow)
	state.stateId = stateData.Id
	return nil
}

func (state *ThirdState) GetName() string {
	return utils.THIRD_STATE_NAME
}
