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

# ====== FROM THE STARTER KIT EXAMPLE =======
#
#
#MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
#VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
#PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
#LDFLAGS:=-ldflags "-X github.com/qiangxue/golang-restful-starter-kit/app.Version=${VERSION}"
#
#default: run
#
#depends:
#	../../../../bin/glide up
#
#test:
#	echo "mode: count" > coverage-all.out
#	$(foreach pkg,$(PACKAGES), \
#		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
#		tail -n +2 coverage.out >> coverage-all.out;)
#
#cover: test
#	go tool cover -html=coverage-all.out
#
#run:
#	go run ${LDFLAGS} server.go
#
#build: clean
#	go build ${LDFLAGS} -a -o server server.go
#
#clean:
#	rm -rf server coverage.out coverage-all.out
