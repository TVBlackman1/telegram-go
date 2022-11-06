package third

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

func (state *ThirdState) joke(msg types.ReceivedMessage) {
	if state.context.JokesInRow >= MAX_JOKES_IN_ROW {
		adviceUserToStopLaught(state, msg)
		return
	}
	joke, err := getRandomJoke(state)
	if err != nil {
		state.queueMessages = append(state.queueMessages, types.Message{
			Text: "Sorry. I do not know interesting jokes...",
		})
	}
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: joke.Text,
	})
	updateJokesInRow(state)
}

func adviceUserToStopLaught(state *ThirdState, msg types.ReceivedMessage) {
	state.queueMessages = append(state.queueMessages, types.Message{
		Text: "Stop laught! Take a break",
	})
}

func getRandomJoke(state *ThirdState) (repository.JokeDbDto, error) {
	jokeCount, err := state.commonContext.Rep.JokeRepository.Count("")
	if err != nil || jokeCount == 0 {
		return repository.JokeDbDto{}, errors.New("zero jokes")
	}
	randomJokeNumber := rand.Int() % jokeCount
	joke, err := state.commonContext.Rep.JokeRepository.GetOne(repository.JokeQuery{
		Offset: randomJokeNumber,
	})
	return joke, err
}

func updateJokesInRow(state *ThirdState) {
	state.context.JokesInRow++
	newContext, _ := json.Marshal(state.context)
	state.commonContext.Rep.JokeRepository.UpdateContext(state.stateId, string(newContext))
}
