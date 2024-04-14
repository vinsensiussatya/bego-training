USE_ENV := $(or $(APP_ENV), local)

ifeq ($(USE_ENV), local)
	ifeq ($(BEGO_ENV), test)
    	include ./params/.env.test
	else
		include ./params/.env
	endif
    export
endif

SHELL         = /bin/sh

APP_NAME      = bego-training
VERSION       = `git describe --always --tags`
GIT_COMMIT    = `git rev-parse HEAD`
GIT_DIRTY     = $(test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE    = `date '+%Y-%m-%d-%H:%M:%S'`
SQUAD         = "bifrost"
CGO_ENABLED   = 0
GOARCH		  = amd64
GOOS		  = $(uname -s)

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make init 				  Init swagger file.'
	@echo '    make build                 Compile the project.'
	@echo '    make run ARGS=             Run with supplied arguments.'
	@echo '    make mocks                 Generate/Update all mock files.'
	@echo '    make test                  Run tests on a compiled project.'
	@echo '    make sec			  		  Run security checks on a compiled project.'
	@echo '    make coverage			  See coverage detail.'
	@echo '    make clean                 Clean the directory tree.'
	@echo '    make get-app-name          Get APP Name.'
	@echo '    make prepare       		  Prepare your branch before make PR.'
	@echo

.PHONY: init
init:
	swag init -g internal/server/swagger_router.go

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -s -X github.com/vinsensiussatya/bego-training/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/vinsensiussatya/bego-training/version.Version=${VERSION} -X github.com/vinsensiussatya/bego-training/version.Environment=${APP_ENV} -X github.com/vinsensiussatya/bego-training/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME} -trimpath .

.PHONY: run
run: init build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: test
test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test -race -coverprofile=coverage.out ./internal/app/...
	go tool cover -func coverage.out

.PHONY: coverage
coverage: test
	@echo "See coverage ${APP_NAME} ${VERSION}"
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}

.PHONY: lint
lint:
	@echo "Check linter with staticcheck"
	staticcheck ./...

.PHONY: sec
sec:
	@echo "Check security issues with gosec"
	gosec -fmt=junit-xml -out=junit.xml -stdout -verbose=text -tests ./...

.PHONY: get-app-name
get-app-name:
	@echo ${APP_NAME}

.PHONY: prepare
prepare: init build mocks test lint sec
	@echo "Your works ready to reviewed! go make the PR"

.PHONY: mocks
mocks:
	mockery --dir ./internal/app/repository --output ./internal/app/repository/mocks --all
	mockery --dir ./internal/app/service --output ./internal/app/service/mocks --all
	mockery --dir ./internal/app/usecase --output ./internal/app/usecase/mocks --all

.PHONY: up
up:
	@docker compose --env-file ./params/.env -p ${APP_NAME} up -d

.PHONY: down
down:
	@docker compose --env-file ./params/.env -p ${APP_NAME} down

# migration
migrate-up:
	migrate -path db/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down

.PHONY: migrate-up migrate-down
