MIGRATION_PATH := ./cmd/migrate/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) force $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run ./cmd/migrate/seed/main.go

.PHONY: run-server
run-server:
	@air

.PHONY: build
build:
	@cd build && docker-compose up -d

.PHONY: clean
clean:
	@cd build && docker-compose down