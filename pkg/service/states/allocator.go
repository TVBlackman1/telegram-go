package states

func GetStateProcessor(stateName string, stateContext *CommonStateContext) (retState UserState) {
	switch stateName {
	case FIRST_STATE_NAME:
		retState = NewFirstState(stateContext)
	case SECOND_STATE_NAME:
		retState = NewSecondState(stateContext)
	case THIRD_STATE_NAME:
		retState = NewThirdState(stateContext)
	}
	return
}
