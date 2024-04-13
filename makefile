.PHONY: run

FILENAME ?= ./assets/qgames.log

run:
	go run ./cmd --path=$(FILENAME)

test:
	go test ./...