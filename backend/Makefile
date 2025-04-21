.PHONY: run run-dev run-prod

# Config
DB_URL="postgres://postgres:postgres@localhost:5439/thichlab_dev?sslmode=disable"
MIGRATE_CMD=migrate -path migrations -database $(DB_URL)

# Default target - run in development mode
run:
	go run main.go

# Run in development mode explicitly
run-dev:
	go run main.go -env dev

# Run in production mode
run-prod:
	go run main.go -env prod

# Apply migration
migrate-up:
	$(MIGRATE_CMD) up

# Rollback migration (1 step)
migrate-down:
	$(MIGRATE_CMD) down 1

# Rollback tất cả migration
migrate-reset:
	$(MIGRATE_CMD) down

# Tạo migration mới (chạy make migrate-create name=migration_name)
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)