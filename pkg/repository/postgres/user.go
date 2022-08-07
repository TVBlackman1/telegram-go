package postgres

import (
	"fmt"
	"os"

	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (rep *UserRepository) Add(interface{}) {

}

func (rep *UserRepository) Remove(interface{}) {

}

func (rep *UserRepository) GetList(interface{}) []repository.UserDbDto {
	query := "select id, name, chat_id, state_id from users"
	var users []repository.UserDbDto
	if err := rep.db.Select(&users, query); err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", err.Error())
	}
	return users
}

func (rep *UserRepository) Edit(interface{}) {

}
