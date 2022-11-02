package types

type Message struct {
	Text      string
	Keyboard  Keyboard
	Media     []Media
	IsCommand bool
}

// TODO integrate message for send: different in keyboard
type MessageForSend struct {
	Message
	Keyboard Keyboard
}

type ChatId int

type Sender struct {
	ChatId ChatId
	Name   string
	Login  string
}

type ReceivedMessage struct {
	Sender  Sender
	Content Message
}
