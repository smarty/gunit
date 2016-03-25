#!/usr/bin/make -f

# This Makefile is an example of what you could feed to scantest's -command flag.

default: test

test: build
	go generate ./...
	go test -v -short ./...

build:
	go build ./...

cover:
	go build
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

install: build
	go install github.com/smartystreets/gunit/gunit
