.PHONY: build run test clean docker-up docker-down

# Build the server binary
build:
	go build -o bin/server ./cmd/server

# Run locally (requires PostgreSQL and Redis running)
run:
	go run ./cmd/server

# Run tests
test:
	go test ./... -v -count=1

# Clean build artifacts
clean:
	rm -rf bin/

# Start all services via Docker
docker-up:
	docker compose up --build -d

# Stop Docker services
docker-down:
	docker compose down -v

# Run database migrations
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

# Download dependencies
deps:
	go mod tidy
	go mod download
