.PHONY: build
build:
	docker compose build
start:
	docker compose up
stop:
	docker compose stop
down:
	docker compose down
restart: down start
init:
	./bin/init.sh
initdb:
	./bin/initdb.sh
backend-shell:
	docker compose exec backend /bin/bash
frontend-shell:
	docker compose exec backend /bin/bash
# backend
ent-add:
	docker compose exec backend go run entgo.io/ent/cmd/ent init ${args}
ent-gen:
	docker compose exec backend go generate ./ent
migrate:
	docker compose exec backend goose.sh up
rollback:
	docker compose exec backend goose.sh down
