.PHONY: run build docker-up docker-down test lint hooks

run:
	go run ./cmd/server

build:
	go build -o bin/app ./cmd/server

docker-up:
	docker compose  up --build

docker-down:
	docker compose down

test:
	go tst ./...

lint:
	pre-commit run --all-files

hooks:
	pre-commit install