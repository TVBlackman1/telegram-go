package postgres

import (
	"github.com/jackc/pgx/v4"
)

type StateRepository struct {
	db *pgx.Conn
}

func NewStateRepository(db *pgx.Conn) *StateRepository {
	return &StateRepository{db}
}

func (rep *StateRepository) Add(interface{}) {

}

func (rep *StateRepository) Remove(interface{}) {

}

func (rep *StateRepository) GetList(interface{}) {
}

func (rep *StateRepository) Edit(interface{}) {

}
