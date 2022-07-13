package make2help

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

const (
	builtInTargetPhony              = ".PHONY"
	builtInTargetSuffixes           = ".SUFFIXES"
	builtInTargetDefault            = ".DEFAULT"
	builtInTargetPrecious           = ".PRECIOUS"
	builtInTargetIntermediate       = ".INTERMEDIATE"
	builtInTargetSecondary          = ".SECONDARY"
	builtInTargetSecondExpansion    = ".SECONDEXPANSION"
	builtInTargetDeleteOnError      = ".DELETE_ON_ERROR"
	builtInTargetIgnore             = ".IGNORE"
	builtInTargetLowResolutionTime  = ".LOW_RESOLUTION_TIME"
	builtInTargetSilent             = ".SILENT"
	builtInTargetExportAllVariables = ".EXPORT_ALL_VARIABLES"
	builtInTargetNotParallel        = ".NOTPARALLEL"
	builtInTargetOneShell           = ".ONESHELL"
	builtInTargetPosix              = ".POSIX"
)

var (
	ruleReg          = regexp.MustCompile(`^([^\s]+)\s*:`)
	isBuiltInTargets = map[string]bool{
		builtInTargetPhony:              true,
		builtInTargetSuffixes:           true,
		builtInTargetDefault:            true,
		builtInTargetPrecious:           true,
		builtInTargetIntermediate:       true,
		builtInTargetSecondary:          true,
		builtInTargetSecondExpansion:    true,
		builtInTargetDeleteOnError:      true,
		builtInTargetIgnore:             true,
		builtInTargetLowResolutionTime:  true,
		builtInTargetSilent:             true,
		builtInTargetExportAllVariables: true,
		builtInTargetNotParallel:        true,
		builtInTargetOneShell:           true,
		builtInTargetPosix:              true,
	}
)

func scan(filepath string) (rules, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, xerrors.Errorf("failed to open file: %w", err)
	}
	r := rules{}
	buf := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "## ") {
			buf = append(buf, line[3:])
			continue
		}

		if matches := ruleReg.FindStringSubmatch(line); len(matches) > 1 {
			target := matches[1]
			if isBuiltInTargets[target] {
				continue
			}
			r[target] = buf
		}
		if len(buf) > 0 {
			buf = []string{}
		}
	}
	if err = sc.Err(); err != nil {
		return nil, xerrors.Errorf("scan failed: %w", err)
	}
	return r, nil
}
