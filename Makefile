SHELL := /bin/bash

APP_NAME := frogshort

# HELP
.PHONY: help
help:
	@echo ""
	@echo "====================== FROGSHORT — MAKE HELP ======================"
	@echo ""
	@echo " Local:"
	@echo "   make run            - run Go service locally (no Docker)"
	@echo ""
	@echo " Compose:"
	@echo "   make compose-build  - build images"
	@echo "   make compose-up     - start stack (db + frogshort)"
	@echo "   make compose-down   - stop stack"
	@echo "   make logs           - follow logs"
	@echo "   make ps             - show containers"
	@echo "   make down-v         - stop + remove volumes"
	@echo ""
	@echo " DB:"
	@echo "   make migrate        - run migrations container once"
	@echo "   make psql           - open psql inside db container"
	@echo ""
	@echo " Go tooling:"
	@echo "   make tidy          - go mod tidy in goShort/"
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
	docker run --rm -p 8070:8070 --name $(APP_NAME) $(APP_NAME):latest

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
