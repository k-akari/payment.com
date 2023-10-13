.PHONY: \
	build \
	lint

build:
	go build -o dist/api cmd/api/*

lint:
	golangci-lint run ./...
