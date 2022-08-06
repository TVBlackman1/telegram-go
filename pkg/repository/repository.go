package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Conn *sqlx.DB
	UserRepository
	StateRepository
}

func (repo *Repository) Close() {
	err := repo.Conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Closing db error: %s", err.Error())
	}
}
