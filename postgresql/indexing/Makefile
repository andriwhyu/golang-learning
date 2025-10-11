MIGRATION_PATH := ./migrations
DOCKER_BUILD_PATH := ./build

.PHONY: migrate-create
migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	@cd $(DOCKER_BUILD_PATH) && docker-compose up -d

.PHONY: run-generator
run-generator:
	@export GENERATED_DATA_COUNT="$(filter-out $@,$(MAKECMDGOALS))" && go run .

.PHONY: clean
clean:
	@cd $(DOCKER_BUILD_PATH) && docker-compose down