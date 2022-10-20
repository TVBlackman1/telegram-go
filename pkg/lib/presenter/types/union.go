package types

type MessageUnion struct {
	Text     string
	Keyboard Keyboard
	Media    []Media
}

type ChatId int

type Sender struct {
	ChatId ChatId
	Name   string
	Login  string
}
type ReceivedMessage struct {
	Sender  Sender
	Content MessageUnion
}
