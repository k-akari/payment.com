.PHONY: \
	build \
	run \
	lint \
	migrate

build:
	go build -o dist/api cmd/api/*

run:
	go run cmd/api/*

lint:
	golangci-lint run ./...

test:
	go test -race ./...

testv:
	go test -race -v ./...

DB_HOST ?= 127.0.0.1
migrate:
	mysqldef --user=root --password=password --port=3306 --host=${DB_HOST} payment < ./db/sql/schema.sql

gen:
	go generate ./...
	$(call fmt)

define fmt
	go fmt ./...
	goimports -w cmd internal
endef
