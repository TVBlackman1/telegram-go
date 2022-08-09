package types

type MessageUnion struct {
	Text     string
	Keyboard Keyboard
	Media    []Media
}

type ChatId int
type ReceivedMessage struct {
	ChatId
	Content MessageUnion
}
