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
	docker compose run -rm server ogen -package ogen -target ogen -clean ./doc/openapi.yaml

.PHONY: apply
apply:
	@until pg_isready -h localhost -p 5432 -U user; do sleep 1; done
	atlas migrate apply \
  --dir "file://server/migrations" \
  --url "postgres://user:password@localhost:5432/db?search_path=public&sslmode=disable"

.PHONY: seed
seed:
	@until pg_isready -h localhost -p 5432 -U user; do sleep 1; done
	PGPASSWORD=$(DB_PASSWORD) psql -f ./server/seed/seed.sql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST)
