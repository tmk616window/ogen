.PHONY: init
init:
	docker compose build	

.PHONY: up
up:
	docker compose up

.PHONY: up-detached
upd:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: gen
gen:
	docker compose exec server ogen -package ogen -target ogen -clean ./doc/openapi.yaml
