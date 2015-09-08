#!/usr/bin/make -f

# This Makefile is an example of what you could feed to scantest's -command flag.

default: test

test: build
	go test ./...

build:
	go build ./...
	go generate ./...

cover:
	go build
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

install:
	go install github.com/smartystreets/gunit/gunit
