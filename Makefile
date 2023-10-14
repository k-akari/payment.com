.PHONY: \
	build \
	lint \
	migrate

build:
	go build -o dist/api cmd/api/*

lint:
	golangci-lint run ./...

test:
	go test -race ./...

testv:
	go test -race -v ./...

migrate:
	mysqldef --user=root --password=password --port=3306 --host=db payment < ./db/sql/schema.sql
