BINARY_NAME=../bin/main.out
MAIN_FILE=../cmd/main.go

PRODUCTION_ENV=production
DEVELOP_ENV=develop
TEST_ENV=test

RED=\033[0;31m
GREEN=\033[0;32m
ORANGE=\033[33m
COMMON_COLOR=\033[0m


notify_about_APP_ENV = $(if $(1),printf "${GREEN}project environment: $(1)${COMMON_COLOR}\n",printf "\${ORANGE}WARNING: APP_ENV is empty${COMMON_COLOR}\n")

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

docker_dev:
	docker run --add-host=host.docker.internal:host-gateway -de APP_ENV=develop bot:0.1
clean:
	go clean
	rm ${BINARY_NAME}