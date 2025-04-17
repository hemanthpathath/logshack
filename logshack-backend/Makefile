# Logshack Backend Makefile

MIGRATE = migrate
DB_URL = postgres://postgres:postgres@localhost:5432/logshack?sslmode=disable

.PHONY: migrate-up migrate-down migrate-new dev stop view-db reset-db test

# Run all migrations
migrate-up:
	$(MIGRATE) -path ./migrations -database "$(DB_URL)" up

# Roll back last migration
migrate-down:
	$(MIGRATE) -path ./migrations -database "$(DB_URL)" down 1

# Create a new migration
migrate-new:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir migrations $$name

# Start containers and run migrations
dev:
	docker compose up -d
	@echo "🕒 Waiting for Postgres to be ready..."
	@sleep 15
	@$(MAKE) migrate-up
	@echo "🚀 Dev environment is ready!"

stop:
	docker compose down
	@echo "🛑 Dev environment stopped!"

# View the database
view-db:
	@echo "🔍 Connecting to database using psql..."
	psql -h localhost -U postgres -d logshack
	
# Fully reset DB and apply migrations
reset-db:
	docker compose down --volumes
	@echo "🧹 Removed DB volume and stopped containers"
	docker compose up -d
	@echo "🕒 Waiting for Postgres to come back..."
	@sleep 5
	@$(MAKE) migrate-up
	@echo "✅ Database reset and migrations applied!"

test:
	@echo "Running tests..."
	go test ./... -v
	@echo "Tests completed!"