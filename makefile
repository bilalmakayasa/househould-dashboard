YQ := yq
MIGRATE := migrate
YAML_FILE := src/config/default.yaml

# Extract database connection information from YAML
DB_HOST := $(shell $(YQ) eval '.database.host' $(YAML_FILE))
DB_PORT := $(shell $(YQ) eval '.database.port' $(YAML_FILE))
DB_USER := $(shell $(YQ) eval '.database.user' $(YAML_FILE))
DB_NAME := $(shell $(YQ) eval '.database.name' $(YAML_FILE))
DB_PASSWORD := $(shell $(YQ) eval '.database.password' $(YAML_FILE))
DB_DRIVER := $(shell $(YQ) eval '.database.driver' $(YAML_FILE))

create-migration:
	@read -p "Enter the migration name: " migration_name; \
	$(MIGRATE) create -ext sql -dir src/database/migrations -seq $$migration_name
	@echo "New migration created successfully."

.PHONY: create-migration

migrate-up:
	$(MIGRATE) -database $(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path src/database/migrations up

.PHONY: migrate-up