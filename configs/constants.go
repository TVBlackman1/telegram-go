package configs

type AplicationEnv string

const (
	APP_ENV = "APP_ENV"
)

const (
	PRODUCTION_ENV AplicationEnv = "production"
	DEVELOP_ENV    AplicationEnv = "develop"
	TEST_ENV       AplicationEnv = "test"
)
