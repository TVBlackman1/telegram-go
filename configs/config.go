package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	TELEGRAM_TOKEN  string
	POSTGRES_USER   string
	POSTGRES_PASS   string
	POSTGRES_DBNAME string
	POSTGRES_PORT   int
}

func LoadConfig(path string) (config Config, err error) {
	appEnviropmentVersion := os.Getenv(APP_ENV)

	configFilename, err := getEnvMapping(appEnviropmentVersion)
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
	fillEmptyEnvValuesByDefault(&config)
	return
}

func fillEmptyEnvValuesByDefault(config *Config) {
	if config.POSTGRES_DBNAME == "" {
		config.POSTGRES_DBNAME = "postgres"
	}
	if config.POSTGRES_PASS == "" {
		config.POSTGRES_PASS = "postgres"
	}
	if config.POSTGRES_USER == "" {
		config.POSTGRES_USER = "postgres"
	}
	if config.POSTGRES_PORT == 0 {
		config.POSTGRES_PORT = 5432
	}
}

func getEnvMapping(appEnviropmentVersion string) (configFilename string, resultError error) {
	switch appEnviropmentVersion {
	case "production":
		configFilename = "production"
	case "develop":
		configFilename = "develop"
	case "test":
		configFilename = "test"
	default:
		errorText := fmt.Sprintf("bad enviropment var: %s", APP_ENV)
		resultError = errors.New(errorText)
	}
	return
}
