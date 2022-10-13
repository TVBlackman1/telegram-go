package configs

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	TELEGRAM_TOKEN  string `mapstructure:"TELEGRAM_TOKEN"`
	POSTGRES_USER   string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASS   string `mapstructure:"POSTGRES_PASS"`
	POSTGRES_DBNAME string `mapstructure:"POSTGRES_DBNAME"`
	POSTGRES_HOST   string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT   int    `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	appEnvironmentVersion := AplicationEnv(os.Getenv(APP_ENV))

	environmentFilename, err := getEnvironmentFilename(appEnvironmentVersion)
	if err != nil {
		return
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(environmentFilename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	setEnvironmentValuesOnFly(&config)

	setEmptyEnvironmentValuesByDefault(&config)
	return
}

func setEnvironmentValuesOnFly(config *Config) (returnErr error) {
	configType := reflect.TypeOf(*config)
	currentConfig := reflect.ValueOf(config).Elem()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		if environmentName, ok := field.Tag.Lookup("mapstructure"); ok {
			foundEnvironment := os.Getenv(environmentName)
			if foundEnvironment != "" {
				currentField := currentConfig.Field(i)
				if currentField.IsValid() && currentField.CanSet() {
					if field.Type.Kind() == reflect.String {
						currentField.SetString(foundEnvironment)
					} else if field.Type.Kind() == reflect.Int {
						stringedInteger, err := strconv.Atoi(foundEnvironment)
						if err != nil {
							returnErr = errors.New("Not corrected int type of config struct")
							return
						}
						currentField.SetInt(int64(stringedInteger))
					}
				} else {
					returnErr = errors.New("config field is not reachable")
					return
				}
			}
		}
	}
	return
}

func setEmptyEnvironmentValuesByDefault(config *Config) {
	println(config.TELEGRAM_TOKEN)
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
		hostOnFly := os.Getenv("POSTGRES_HOST")
		// TODO simplify, auto loading from env space
		if hostOnFly != "" {
			config.POSTGRES_HOST = hostOnFly
		} else {
			config.POSTGRES_HOST = "0.0.0.0"
		}
	}
}

func getEnvironmentFilename(appEnvironmentVersion AplicationEnv) (configFilename string, resultError error) {
	switch appEnvironmentVersion {
	case PRODUCTION_ENV:
		configFilename = "production"
	case DEVELOP_ENV:
		configFilename = "develop"
	case TEST_ENV:
		configFilename = "test"
	default:
		errorText := fmt.Sprintf("bad environment var: %s", APP_ENV)
		resultError = errors.New(errorText)
	}
	return
}
