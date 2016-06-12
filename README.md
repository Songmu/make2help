make2help
=======

[![Build Status](https://travis-ci.org/Songmu/make2help.png?branch=master)][travis]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/make2help?status.svg)](godoc)

[travis]: https://travis-ci.org/Songmu/make2help
[coveralls]: https://coveralls.io/r/Songmu/make2help?branch=master
[license]: https://github.com/Songmu/make2help/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/make2help

## Description

Utility for self-documented Makefile

It scans Makefiles and shows rules with documents. It considers the comment line started with
double hash (`## `) just before a rule is written as document of the rule.

## Installation

    % go get github.com/Songmu/make2help/cmd/make2help

## Synopsis

    % make2help
    Available rules:

    cover              Take coverage
    deps               Install dependencies
    dev-deps           Install dependencies
    help               Show help
    lint               Lint
    release            Release the binaries
    test               Run tests

## Options

```
-all                display all rules in the Makefiles
```

## Example

With defining `help` target in Makefile and setting it to `.DEFAULT_GOAL`, you can see
help messages only command `make`.

    % cat Makefile
    .DEFAULT_GOAL := help
    
    ## Run tests
    test: deps
        go test ./...
    
    ## Install dependencies
    deps:
        go get -d -v -t ./...
    
    ## Show help
    help:
        @make2help $(MAKEFILE_LIST)
    
    .PHONY: test deps help

    % make
    Available rules:

    deps               Install dependencies
    help               Show help
    test               Run tests

## Author

[Songmu](https://github.com/Songmu)
