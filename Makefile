SHELL := /bin/bash

APP_NAME     := frogshort
CMD_DIR      := $(SERVICE_DIR)/cmd
BINARY       := $(APP_NAME)

# HELP
.PHONY: help
help:
	@echo ""
	@echo "====================== FROGSHORT — MAKE HELP ======================"
	@echo ""
	@echo " Main commands:"
	@echo "   make run           - run Go service locally (no Docker)"
	@echo ""
	@echo " Docker / Compose:"
	@echo "   make compose-up    - docker compose up --build (app + Postgres)"
	@echo "   make compose-down  - docker compose down"
	@echo "   make re            - restart stack (down + up)"
	@echo "   make logs          - docker compose logs -f --tail=200"
	@echo "   make ps            - docker compose ps"
	@echo "   make down-v        - WARNING: down + remove volumes (DB data!)"
	@echo ""
	@echo " Go tooling:"
	@echo "   make tidy          - go mod tidy in goShort/"
	@echo ""
	@echo "==================================================================="
	@echo ""

# GO LOCAL
.PHONY: run build clean \
				tidy fmt
run:
	go run ./cmd/$(APP_NAME)/main.go

tidy:
	go mod tidy

# DOCKER / COMPOSE
.PHONY: docker-build docker-run docker-stop \
        compose-up compose-down \
				logs ps down-v \
				re
docker-build:
	docker compose build

docker-run:
	docker run --rm -p 8080:8080 --name $(APP_NAME) $(APP_NAME):latest

docker-stop:
	- docker stop $(APP_NAME)

compose-up:
	docker compose up --build

compose-down:
	docker compose down

logs:
	docker compose logs -f --tail=200

ps:
	docker compose ps

down-v:
	docker compose down -v

re: compose-down compose-up

# DATABASE
.PHONY: db 
db:
	docker compose exec db psql "postgres://dev:dev@127.0.0.1:5432/dev_db?options=-c%20search_path%3Dfrogshort"
