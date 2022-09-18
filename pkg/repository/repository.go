package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Conn            *sqlx.DB
	UserRepository  UserRepository
	StateRepository StateRepository
}

func (repo *Repository) Close() {
	err := repo.Conn.Close()
	if err != nil {
		log.Fatalf("Closing db error: %s", err.Error())
	}
}
