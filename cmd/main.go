package main

import (
	"log"

	"github.com/TVBlackman1/telegram-go/configs"
	"github.com/TVBlackman1/telegram-go/pkg"
)

func main() {
	config, err := configs.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	pkg.RunTelegramBot(config.TELEGRAM_TOKEN)
}
