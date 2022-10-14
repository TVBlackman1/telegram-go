package repository

import (
	"github.com/google/uuid"
)

type StateRepository interface {
	Add(query CreateStateDto) (uuid.UUID, error)
	Remove(interface{})
	Edit(interface{})
	GetList(interface{})
}

type CreateStateDto struct {
	Id      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Context string    `db:"context"`
}

type StateDbDto struct {
	Id      uuid.UUID
	Name    string
	Context string
}

const (
	STATES_TABLENAME = "states"
)
