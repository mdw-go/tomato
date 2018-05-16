#!/usr/bin/make -f

test:
	go build ./...
	go test -v -short ./...

install:
	go install ./cmd/...

docs:
	go get -u github.com/robertkrimen/godocdown/godocdown
	godocdown > README.md
