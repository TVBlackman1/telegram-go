package states

func DefineState(stateName string) UserState {
	return new(FirstState)
}
