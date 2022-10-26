package postgres

import (
	"github.com/TVBlackman1/telegram-go/pkg/repository"
)

func NewRepository(config PostgresConfig) (*repository.Repository, error) {
	conn, err := Connect(config)
	if err != nil {
		return nil, err
	}
	repo := new(repository.Repository)
	repo.Conn = conn
	repo.UserRepository = NewUserRepository(repo.Conn)
	repo.StateRepository = NewStateRepository(repo.Conn)
	repo.JokeRepository = NewJokeRepository(repo.Conn)

	return repo, nil
}
