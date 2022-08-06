package postgres

import (
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/jackc/pgx/v4"
)

type PostgresRepository struct {
	conn            *pgx.Conn
	UserRepository  repository.UserRepository
	StateRepository repository.StateRepository
}

func NewRepository(config PostgresConfig) (*PostgresRepository, error) {
	conn, err := Connect(config)
	if err != nil {
		return nil, err
	}
	repo := new(PostgresRepository)
	repo.conn = conn
	repo.UserRepository = NewUserRepository(repo.conn)
	repo.StateRepository = NewStateRepository(repo.conn)

	return repo, nil
}
