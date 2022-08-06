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
	POSTGRES_HOST   string
	POSTGRES_PORT   int
}

func LoadConfig(path string) (config Config, err error) {
	appEnviropmentVersion := AplicationEnv(os.Getenv(APP_ENV))

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
	if config.POSTGRES_HOST == "" {
		config.POSTGRES_HOST = "0.0.0.0"
	}
}

func getEnvMapping(appEnviropmentVersion AplicationEnv) (configFilename string, resultError error) {
	switch appEnviropmentVersion {
	case PRODUCTION_ENV:
		configFilename = "production"
	case DEVELOP_ENV:
		configFilename = "develop"
	case TEST_ENV:
		configFilename = "test"
	default:
		errorText := fmt.Sprintf("bad enviropment var: %s", APP_ENV)
		resultError = errors.New(errorText)
	}
	return
}
