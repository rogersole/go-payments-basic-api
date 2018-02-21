.PHONY: image build run test

image:
	docker build -t payments-basic-api .

build:
	go build -o payments-basic-api .

run:
	make build
	docker-compose up -d postgres
	PAYMENTS_DSN=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable ./payments-basic-api

test:
	docker-compose up -d postgres
	sleep 3
	PAYMENTS_DSN=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable go test ./...
