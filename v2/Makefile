#!/usr/bin/make -f

test: fmt
	go test -timeout=1s -race -covermode=atomic ./...
	go test -timeout=1s -race -covermode=atomic -short ./...

fmt:
	go mod tidy && go fmt ./...

compile:
	go build ./...

build: test compile

.PHONY: fmt test compile build
