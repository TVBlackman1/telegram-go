package states

import (
	"fmt"
	"math/rand"

	"github.com/TVBlackman1/telegram-go/pkg/lib"
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

const THIRD_STATE_NAME = "Third state"

type ThirdState struct {
	commonContext *CommonStateContext
	jokeToAnswer  bool
	queueMessages []types.Message
	autoMessages  []notifier.NotifierContext
}

func NewThirdState(context *CommonStateContext) *ThirdState {
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
		newState := StateOnFlyDto{SECOND_STATE_NAME, "{}"}
		stateSwitcher.TransferToNewStateByChatId(chatId, newState)
		state.queueMessages = append(state.queueMessages, types.Message{
			Text: "Transfer to second state",
		})
	}
	if msg.Content.Text == "joke" {
		jokeCount, _ := state.commonContext.rep.JokeRepository.Count("")
		if jokeCount == 0 {
			state.queueMessages = append(state.queueMessages, types.Message{
				Text: fmt.Sprint("Sorry. I do not know interesting jokes..."),
			})
		} else {
			randomJokeNumber := rand.Int() % jokeCount
			joke, _ := state.commonContext.rep.JokeRepository.GetOne(repository.JokeQuery{
				Offset: randomJokeNumber,
			})
			state.queueMessages = append(state.queueMessages, types.Message{
				Text: fmt.Sprintf("%s", joke.Text),
			})
		}
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

func (state *ThirdState) SetContext(msg types.ReceivedMessage, context interface{}) error {
	return nil
}

func (state *ThirdState) GetName() string {
	return THIRD_STATE_NAME
}
