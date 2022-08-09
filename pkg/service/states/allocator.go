package states

type StateId int

const (
	FIRST_STATE = iota
	SECOND_STATE
)

func DefineState(stateId StateId) (retState UserState) {
	switch stateId {
	case FIRST_STATE:
		retState = new(FirstState)
	case SECOND_STATE:
		retState = new(SecondState)
	}
	return
}
