package main

import (
	"os"

	"github.com/Songmu/make2help"
)

func main() {
	os.Exit((&make2help.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
