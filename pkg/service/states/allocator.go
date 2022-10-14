package states

func DefineState(stateName string) (retState UserState) {
	switch stateName {
	case FIRST_STATE_NAME:
		retState = new(FirstState)
	case SECOND_STATE_NAME:
		retState = new(SecondState)
	}
	return
}
