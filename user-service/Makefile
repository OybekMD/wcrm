# Define variables
CURRENT_DIR := $(shell pwd)
APP := comment-service
APP_CMD_DIR := cmd
DB_URL := "postgres://postgres:ebot@localhost:5432/wcrmdb?sslmode=disable"

# Display available commands
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Run the service
run: ## Run the service
	go run $(APP_CMD_DIR)/main.go

# Build the application binary
build: ## Build the application binary
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o $(CURRENT_DIR)/bin/$(APP) $(APP_CMD_DIR)/main.go

# Generate protobuf files
proto-gen: ## Generate protobuf files
	./scripts/gen-proto.sh

# Apply database migrations
migrate-up: ## Apply database migrations
	migrate -path migrations -database $(DB_URL) -verbose up

# Rollback database migrations
migrate-down: ## Rollback database migrations
	migrate -path migrations -database $(DB_URL) -verbose down

# Create a new migration file
migrate-file: ## Create a new migration file
	migrate create -ext sql -dir migrations/ -seq $(name)

# Apply all unapplied migrations forcefully
migrate-dirty: ## Apply all unapplied migrations forcefully
	migrate -path ./migrations/ -database $(DB_URL) force 1

# Check the current migration version
migration-version: ## Check the current migration version
	migrate -database $(DB_URL) -path migrations version 

# Create a new migration
create-migration: ## Create a new migration
	migrate create -ext sql -dir migrations -seq "$(name)"
