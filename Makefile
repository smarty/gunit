#!/usr/bin/make -f

test-dev: fmt
	go test -timeout=1s -race -count=1 -covermode=atomic ./... && \
	  echo "---- Running tests a second time with -short ----" && \
	  go test -timeout=1s -race -count=1 -covermode=atomic -short ./...

fmt:
	go fmt ./...

test:
	go test -timeout=1s -race -covermode=atomic .

compile:
	go build ./...

build: test compile

.PHONY: test-dev fmt test compile build
