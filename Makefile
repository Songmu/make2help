VERSION = $(shell godzil show-version)
CURRENT_VERSION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X github.com/Songmu/make2help.revision=$(CURRENT_REVISION)"
u := $(if $(update),-u)

.DEFAULT_GOAL := help
export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get ${u} -d -v
	go mod tidy

## Run tests
.PHONY: test
test:
	go test

## Install dependencies
.PHONY: devel-deps
devel-deps:
	go install github.com/Songmu/godzil/cmd/godzil@latest
	go install github.com/tcnksm/ghr@latest

bin/%: cmd/%/main.go
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## Bump version
.PHONY: bump
bump: devel-deps
	godzil release

CREDITS: deps devel-deps go.sum
	godzil credits -w

## Cross build
.PHONY: crossbuild
crossbuild: CREDITS
	godzil crossbuild -pv=v$(VERSION) -build-ldflags=$(BUILD_LDFLAGS) \
	  -os=linux,darwin,freebsd,windows -d=./dist/v$(VERSION) ./cmd/*

## Upload
.PHONY: upload
upload:
	ghr v$(VERSION) dist/v$(VERSION)

## Release the binaries
.PHONY: release
release: crossbuild upload

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
