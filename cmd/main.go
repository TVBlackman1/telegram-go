package main

import (
	"log"

	"github.com/TVBlackman1/telegram-go/configs"
	"github.com/TVBlackman1/telegram-go/pkg/notifier"
	repo "github.com/TVBlackman1/telegram-go/pkg/repository/postgres"
	"github.com/TVBlackman1/telegram-go/pkg/router"
	"github.com/TVBlackman1/telegram-go/pkg/service"
	telegramlistener "github.com/TVBlackman1/telegram-go/pkg/telegram-listener"
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
	if err != nil {
		log.Fatalf("repo err: %s", err.Error())
	}
	defer repo.Close()

	notifier := notifier.NewSystemNotifier()
	// TODO check how to use notifier just in user service
	userService := service.NewUserService(repo, notifier)
	router := router.NewRouter(userService)
	bot := telegramlistener.NewTelegramBot(config.TELEGRAM_TOKEN, router, notifier)
	// go func() {
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		bot.NotifyUser(537172348)
	// 	}
	// }()
	bot.Run()
}
