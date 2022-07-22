BINARY_NAME=main.out
MAIN_FILE=../cmd/main.go

PRODUCTION_ENV=production
DEVELOP_ENV=develop
TEST_ENV=test


notify_about_APP_ENV = $(if $(1),echo "\033[0;32mproject enviropment: $(1)\033[0m",echo "\033[33mWARNING: APP_ENV is empty\033[0m")

run_production: export APP_ENV=${PRODUCTION_ENV}
run_develop: export APP_ENV=${DEVELOP_ENV}
run_test: export APP_ENV=${TEST_ENV}

run_production run_develop run_test: run
	@echo ${APP_ENV}

build:
	go build -mod=vendor -o ${BINARY_NAME} ${MAIN_FILE}

run: build
	@$(call notify_about_APP_ENV,${APP_ENV})
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}