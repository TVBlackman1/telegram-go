package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db}
}

func (rep *UserRepository) Add(interface{}) {

}

func (rep *UserRepository) Remove(interface{}) {

}

func (rep *UserRepository) GetList(interface{}) {
	var name string
	rep.db.QueryRow(context.Background(), "select name from users").Scan(&name)
	fmt.Println(name)
}

func (rep *UserRepository) Edit(interface{}) {

}
