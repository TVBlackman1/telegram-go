package postgres

import "github.com/jmoiron/sqlx"

type StateRepository struct {
	db *sqlx.DB
}

func NewStateRepository(db *sqlx.DB) *StateRepository {
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
