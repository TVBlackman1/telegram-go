package repository

type UserRepository interface {
	Add(interface{})
	Remove(interface{})
	Edit(interface{})
	GetList(interface{})
}
