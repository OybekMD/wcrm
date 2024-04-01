# Define variables
CURRENT_DIR := $(shell pwd)
APP := comment-service
APP_CMD_DIR := cmd
DB_URL := "postgres://postgres:ebot@localhost:5432/wcrmdb?sslmode=disable"

help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the service
	go run $(APP_CMD_DIR)/main.go

build: ## Build the application binary
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o $(CURRENT_DIR)/bin/$(APP) $(APP_CMD_DIR)/main.go

proto-gen: ## Generate protobuf files
	./scripts/gen-proto.sh