package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type PostgresConfig struct {
	Dbname   string
	Host     string
	Port     int
	User     string
	Password string
}

func Connect(config PostgresConfig) (*pgx.Conn, error) {
	dbConnectUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	conn, err := pgx.Connect(context.Background(), dbConnectUrl)
	if err != nil {
		errorText := fmt.Sprintf("Unable to connect to database: %v\n", err)
		return nil, errors.New(errorText)
	}
	fmt.Printf("Connected to database '%s' on %s:%d\n", config.Dbname, config.Host, config.Port)
	// defer conn.Close(context.Background())
	return conn, nil
}
