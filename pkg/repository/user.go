package repository

type UserRepository interface {
	Add(interface{})
	Remove(interface{})
	Edit(interface{})
	GetList(interface{}) []UserDbDto
}

type UserDbDto struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	ChatId  int    `db:"chat_id"`
	StateId string `db:"state_id"`
}
