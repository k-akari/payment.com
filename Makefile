.PHONY: \
	build

build:
	go build -o dist/api cmd/api/*
