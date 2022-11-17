#!/usr/bin/make -f


all: creatavisor test

creatavisor:
	go build -mod=readonly ./cmd/creatavisor

test:
	go test -mod=readonly -race ./...

.PHONY: all creatavisor test
