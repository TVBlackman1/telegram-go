package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type UserDbDto struct {
	Name string
}

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

func (rep *UserRepository) GetList(interface{}) {
	var name string
	rep.db.QueryRow("select name from users").Scan(&name)
	query := "select name from users"
	rows, err := rep.db.Queryx(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", query)
	}
	for rows.Next() {
		var user UserDbDto
		rows.StructScan(&user)
		fmt.Println(user.Name)
	}
}

func (rep *UserRepository) Edit(interface{}) {

}
