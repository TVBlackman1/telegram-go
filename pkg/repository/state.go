package repository

type StateRepository interface {
	Add(interface{})
	Remove(interface{})
	Edit(interface{})
	GetList(interface{})
}
