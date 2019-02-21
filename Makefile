VERSION = $(shell godzil show-version)
CURRENT_VERSION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X github.com/Songmu/make2help.revision=$(CURRENT_REVISION)"
ifdef update
  u=-u
endif

.DEFAULT_GOAL := help
export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get ${u} -d -v

## Run tests
.PHONY: test
test: deps
	go test

## Install dependencies
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get ${u} \
	  golang.org/x/lint/golint            \
	  github.com/mattn/goveralls          \
	  github.com/Songmu/godzil/cmd/godzil \
	  github.com/Songmu/goxz/cmd/goxz     \
	  github.com/tcnksm/ghr

## Lint
.PHONY: lint
lint: devel-deps
	go vet
	golint -set_exit_status

## Take coverage
.PHONY: cover
cover: devel-deps
	goveralls

bin/%: cmd/%/main.go
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## Bump version
.PHONY: bump
bump: devel-deps
	godzil release

## Cross build
.PHONY: crossbuild
crossbuild:
	goxz -pv=v$(VESION) -build-ldflags=$(BUILD_LDFLAGS) \
	  -os=linux,darwin,freebsd,windows -d=./dist/v$(VERSION) ./cmd/*

## Upload
.PHONY: upload
upload:
	ghr v$(VERSION) dist/v$(VERSION)

## Release the binaries
.PHONY: release
release: bump crossbuild upload

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
