.PHONY: run test lint migrate-up migrate-down

run:
	go run ./cmd/server/...

test:
	go test -race ./...

lint:
	golangci-lint run ./...

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down 1