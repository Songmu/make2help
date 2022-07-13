make2help
=======

[![Test Status](https://github.com/Songmu/make2help/workflows/test/badge.svg?branch=main)][actions]
[![MIT License](https://img.shields.io/github/license/Songmu/make2help)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Songmu/make2help)][PkgGoDev]

[actions]: https://github.com/Songmu/make2help/actions?workflow=test
[license]: https://github.com/Songmu/make2help/blob/master/LICENSE
[PkgGoDev]: https://pkg.go.dev/github.com/Songmu/make2help

## Description

Utility for self-documented Makefile

It scans Makefiles and shows rules with documents. It considers the comment line started with
double hash (`## `) just before a rule is written as document of the rule.

## Installation

Binaries are available.

https://github.com/Songmu/make2help/releases

You can also `go install`.

    % go install github.com/Songmu/make2help/cmd/make2help@latest

## Synopsis

    % make2help
    cover:             Take coverage
    deps:              Install dependencies
    dev-deps:          Install dependencies
    help:              Show help
    lint:              Lint
    release:           Release the binaries
    test:              Run tests

## Options

```
-all                display all rules in the Makefiles
```

## Example

With defining `help` target in Makefile and setting it to `.DEFAULT_GOAL`, you can see
help messages just type `make`.

```Makefile
.DEFAULT_GOAL := help

## Run tests
test: deps
    go test ./...

## Install dependencies
deps:
    go get -d -v ./...

## Show help
help:
    @make2help $(MAKEFILE_LIST)

.PHONY: test deps help
```

```Shell
% make
deps:              Install dependencies
help:              Show help
test:              Run tests
```

## Author

[Songmu](https://github.com/Songmu)
