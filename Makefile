.DEFAULT_GOAL := help

## Run tests
test: deps
	go test ./...

## Install dependencies
deps:
	go get -d -v -t ./...

## Install dependencies
dev-deps: deps
	go get github.com/golang/lint/golint
	go get github.com/mattn/goveralls
	go get github.com/motemen/gobump/cmd/gobump
	go get github.com/Songmu/ghch/cmd/ghch
	go get github.com/laher/goxc
	go get github.com/tcnksm/ghr

## Lint
lint: dev-deps
	go vet ./...
	golint -set_exit_status ./...

## Take coverage
cover: dev-deps
	goveralls

## Release the binaries
release:
	_tools/releng

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test dev-deps deps lint cover release help
