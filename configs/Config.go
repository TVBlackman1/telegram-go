package configs

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	TELEGRAM_TOKEN string
}

func LoadConfig(path string) (config Config, err error) {
	APP_ENV := os.Getenv("APP_ENV")

	configFilename, err := getEnvMapping(APP_ENV)
	if err != nil {
		return
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(configFilename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func getEnvMapping(APP_ENV string) (configFilename string, resultError error) {
	switch APP_ENV {
	case "production":
		configFilename = "production"
	case "develop":
		configFilename = "develop"
	case "test":
		configFilename = "test"
	default:
		resultError = errors.New("bad enviropment var: APP_ENV")
	}
	return
}
