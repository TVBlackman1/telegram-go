package presenter

type MessageUnion struct {
	Text     string
	Keyboard Keyboard
	Media    []Media
}
