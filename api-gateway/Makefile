CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd
DB_URL := "postgres://postgres:ebot@localhost:5432/wcrmdb?sslmode=disable"

help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run currently service
	go run cmd/main.go

test: ## Tests all test files
	go test -v ./...

jprq: ## Run api in JPRQ server
	jprq http 8090 -s wcrm

build: ## Build Unknown
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen: ## Generate protocol buffers
	./scripts/gen-proto.sh	${CURRENT_DIR}

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-gen: ## Generate swagger for api
	swag init -g api/router.go -o api/docs

migrate-up: ## Apply database migrations
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down: ## Rollback database migrations
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-file: ## Create a new migration file
	migrate create -ext sql -dir migrations/ -seq $(name)

migrate-dirty: ## Apply all unapplied migrations forcefully
	migrate -path ./migrations/ -database $(DB_URL) force 1

migration-version: ## Check the current migration version
	migrate -database $(DB_URL) -path migrations version 

create-migration: ## Create a new migration
	migrate create -ext sql -dir migrations -seq "$(name)"