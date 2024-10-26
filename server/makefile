DB_USER = user
DB_NAME = db
DB_HOST = localhost
DB_PASSWORD = password

.PHONY: init
init:
	@docker compose build --no-cache
	@make upd
	@make apply
	@make seed

.PHONY: up
up:
	docker compose up

.PHONY: upd
upd:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: gen
gen:
	docker compose run --rm server ogen -package ogen -target ogen -clean ./doc/openapi.yaml

.PHONY: apply
apply:
	@until pg_isready -h $(DB_HOST) -p 5432 -U $(DB_USER); do sleep 1; done
	atlas migrate apply \
  --dir "file://migrations" \
  --url "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_NAME)?search_path=public&sslmode=disable"

.PHONY: migrate_diff
migrate_diff:
	@until pg_isready -h $(DB_HOST) -p 5432 -U $(DB_USER); do sleep 1; done
	atlas migrate diff migration_name \
	  --dir "file://migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/15/test?search_path=public"

.PHONY: seed
seed:
	@until pg_isready -h $(DB_HOST) -p 5432 -U $(DB_USER); do sleep 1; done
	PGPASSWORD=$(DB_PASSWORD) psql -f ./seed/seed.sql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST)
