.PHONY: image build run test

image:
	docker build -t payments-basic-api .

build:
	go build -o payments-basic-api ./cmd/api/main.go

run:
	make build
	docker-compose up -d postgres
	POSTGRESQL_URL=localhost:6379 \
	./payments-basic-api

test:
	docker-compose up -d postgres
	sleep 3
	go test ./...                # test will launch ml-data-api passing all the needed ENV VARS
