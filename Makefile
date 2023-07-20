#!/usr/bin/make -f

test-dev: fmt
	go test -timeout=1s -count=1 ./... && \
	  echo "---- Running tests a second time with -short ----" && \
	  go test -timeout=1s -count=1 -short ./...

fmt:
	go fmt ./...

test:
	go test -timeout=1s -race -covermode=atomic .

compile:
	go build ./...

build: test compile

.PHONY: test compile build
