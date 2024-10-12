.PHONY: init
init:
	docker compose build	

.PHONY: up
run:
	docker compose up

.PHONY: upd
run:
	docker compose up -d

.PHONY: stop
stop:
	docker compose down

.PHONY: gen
gen:
	docker compose exec server ogen -package ogen -target ogen -clean ./doc/openapi.yaml
