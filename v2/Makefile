#!/usr/bin/make -f

test: fmt
	GORACE="atexit_sleep_ms=50" go test -timeout=1s -race -count=1 -covermode=atomic ./...

fmt:
	go mod tidy && go fmt ./...

compile:
	go build ./...

build: test compile

.PHONY: fmt test compile build
