# Simple Makefile for a Go project
#
# Compile e rode os testes.
# ```bash
# make all
# ```

# Compile a aplicação
# ```bash
# make build
# ```

# Rode a aplicação
# ```bash
# make run
# ```

# Cria o banco de dados por um _container_ docker:
# ```bash
# make docker-run
# ```

# Desativa o container:
# ```bash
# make docker-down
# ```

# Testes de integração no Banco de Dados:
# ```bash
# make itest
# ```

# Live reload aplicação:
# ```bash
# make watch
# ```

# Rode os testes:
# ```bash
# make test
# ```

# Remova o executavél
# ```bash
# make clean
# ```
# Build the application
all: build test

build:
	@echo "Building..."


	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build -d database 2>/dev/null; then \
		echo "Running Docker Compose" ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build database; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down database 2>/dev/null; then \
		echo "Shutdown Docker Compose" ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

docs:
	swag init --dir cmd/api,internal --parseDependency --parseInternal

help:
	@echo "Usage: make command"
	@echo "Execute commands quick commands on the project with the help of make.\n"
	@echo "COMMANDS:"
	@echo "  all         - Build the application and run unit tests (default target)"
	@echo "  build       - Compile the Go application (outputs ./main)"
	@echo "  run         - Run the application using 'go run cmd/api/main.go'"
	@echo "  docker-run  - Create and start the database container (Docker Compose)"
	@echo "  docker-down - Stop and remove the database container (Docker Compose)"
	@echo "  test        - Run all unit tests (go test ./... -v)"
	@echo "  itest       - Run integration tests for the database package (internal/database)"
	@echo "  watch       - Start live-reload using 'air' (offers to install if missing)"
	@echo "  clean       - Remove the compiled binary (rm -f main)"
	@echo "  docs        - Generate Swagger docs using 'swag init'"
	@echo "  help        - Show this help message"

.PHONY: all build run test clean watch docker-run docker-down itest docs help
