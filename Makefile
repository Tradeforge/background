include ./Makefile.vars

GO := $(shell which go)

.PHONY:
	run  \
	test \
	build

all: fmt vet build

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

run: fmt vet
	$(GO) run ./cmd

test: lint
	$(GO) test ./... -cover

lint:
	golangci-lint run
