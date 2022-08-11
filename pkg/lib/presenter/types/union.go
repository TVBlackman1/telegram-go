package types

type MessageUnion struct {
	Text     string
	Keyboard Keyboard
	Media    []Media
}

type ChatId int

type Sender struct {
	ChatId
	Name  string
	Login string
}
type ReceivedMessage struct {
	Sender
	Content MessageUnion
}
