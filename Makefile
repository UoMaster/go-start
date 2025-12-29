# Import .env file
ifneq (,$(wildcard ./.env))
		include .env
		export $(shell sed 's/=.*//' .env)
endif

# Variables
CONTAINER_NAME=${APP_NAME}-app
POSTGRES_CONTAINER_NAME=${APP_NAME}-db

# Commands
dep: 
	go mod tidy

run: 
	go run cmd/main.go

build: 
	go build -o main cmd/main.go

run-build: build
	./main

test:
	go test -v ./...

# Postgres commands
container-postgres:
	docker exec -it ${POSTGRES_CONTAINER_NAME} /bin/sh

create-db:
	docker exec -it ${POSTGRES_CONTAINER_NAME} /bin/sh -c "createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}"

init-uuid:
	docker exec -it ${POSTGRES_CONTAINER_NAME} /bin/sh -c "psql -U ${DB_USER} -d ${DB_NAME} -c 'CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";'"

# Docker commands
init-docker:
	docker compose up -d --build

up: 
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

container-go:
	docker exec -it ${CONTAINER_NAME} /bin/sh
