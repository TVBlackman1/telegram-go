package states

import "errors"

type StateId int

const (
	FIRST_STATE = iota
	SECOND_STATE
)

func (stateId StateId) String() (string, error) {
	switch stateId {
	case FIRST_STATE:
		return "FirstState", nil
	case SECOND_STATE:
		return "SecondState", nil
	default:
		return "", errors.New("stateId validation error")
	}
}

func DefineState(stateId StateId) (retState UserState) {
	switch stateId {
	case FIRST_STATE:
		retState = new(FirstState)
	case SECOND_STATE:
		retState = new(SecondState)
	}
	return
}
