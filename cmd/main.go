package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TVBlackman1/telegram-go/configs"
	"github.com/TVBlackman1/telegram-go/pkg"
	"github.com/TVBlackman1/telegram-go/pkg/repository"
	repo "github.com/TVBlackman1/telegram-go/pkg/repository/postgres"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
)

func main() {
	config, err := configs.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig := repo.PostgresConfig{
		Dbname:   config.POSTGRES_DBNAME,
		Host:     config.POSTGRES_HOST,
		Port:     config.POSTGRES_PORT,
		User:     config.POSTGRES_USER,
		Password: config.POSTGRES_PASS,
	}
	repo, err := repo.NewRepository(dbConfig)
	defer repo.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "repo err: %s", err.Error())
		os.Exit(1)
	}
	repo.UserRepository.GetList(repository.UserListQuery{
		Pagination: utils.QueryPagination{
			Size: 20,
			Page: 1,
		},
	})
	pkg.RunTelegramBot(config.TELEGRAM_TOKEN)
}
