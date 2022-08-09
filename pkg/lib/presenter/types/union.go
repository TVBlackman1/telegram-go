package types

type MessageUnion struct {
	Text     string
	Keyboard Keyboard
	Media    []Media
}
