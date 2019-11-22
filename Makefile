#!/usr/bin/make -f

test-dev: fmt
	go test -timeout=1s -count=1 ./...

fmt:
	go fmt ./...

test:
	go test -timeout=1s -race -coverprofile=coverage.txt -covermode=atomic .

compile:
	go build ./...

build: test compile

.PHONY: test compile build
