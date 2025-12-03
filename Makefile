SHELL := /bin/bash

APP_NAME     := frogshort
SERVICE_DIR  := goShort
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
	@echo "   make build         - build Go binary locally"
	@echo "   make clean         - remove local binary"
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
	@echo "   make fmt           - go fmt ./... in goShort/"
	@echo ""
	@echo "==================================================================="
	@echo ""

# GO LOCAL
.PHONY: run build clean \
				tidy fmt
run:
	cd $(SERVICE_DIR) && go run ./cmd/main.go

build:
	cd $(SERVICE_DIR) && go build -o ../$(BINARY) ./cmd/main.go

clean:
	rm -f $(BINARY)

tidy:
	cd $(SERVICE_DIR) && go mod tidy

fmt:
	cd $(SERVICE_DIR) && go fmt ./...

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
	docker compose exec db psql -U "$$DB_USER" "$$DB_NAME"
