package presenter

type MessageUnion struct {
	text     string
	keyboard Keyboard
	media    []Media
}
