package make2help

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

const version = "0.2.0"

var revision = "HEAD"

const (
	exitCodeOK = iota
	exitCodeParseFlagError
	exitCodeErr
)

// CLI is struct for command line tool
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run the make2help
func (cli *CLI) Run(argv []string) int {
	argv, isHelp, all := parseFlags(argv)
	if isHelp {
		fmt.Fprintln(cli.ErrStream, help())
		return exitCodeOK
	}
	if len(argv) < 1 {
		argv = []string{"Makefile"}
	}
	colorful := false
	if w, ok := cli.OutStream.(*os.File); ok {
		colorful = isatty.IsTerminal(w.Fd())
		cli.OutStream = colorable.NewColorable(w)
	}
	r := rules{}
	for _, f := range argv {
		tmpRule, err := scan(f)
		if err != nil {
			fmt.Fprintln(cli.ErrStream, err)
			return exitCodeErr
		}
		r = r.merge(tmpRule)
	}
	fmt.Fprint(cli.OutStream, r.string(all, colorful))
	return exitCodeOK
}

var (
	helpReg = regexp.MustCompile(`^--?h(?:elp)?$`)
	allReg  = regexp.MustCompile(`^--?all$`)
)

func parseFlags(argv []string) (restArgv []string, isHelp, isAll bool) {
	for _, v := range argv {
		if helpReg.MatchString(v) {
			isHelp = true
			return
		}
		if allReg.MatchString(v) {
			isAll = true
			continue
		}
		restArgv = append(restArgv, v)
	}
	return
}

func help() string {
	return fmt.Sprintf(`Usage:
  $ make2help [Makefiles]

Utility for self-documented Makefile

It shows rules in Makefiles with documents.

Options:
  -all          display all rules in the Makefiles

Version: %s`, version)
}
