package make2help

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

var (
	ruleReg          = regexp.MustCompile(`^(\S+):`)
	builtinTargetReg = regexp.MustCompile(`^\.[A-Z_]{5,}`) // ex. ".PHONY"
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
			if target == ".PHONY" {
				continue
			}
			if !builtinTargetReg.MatchString(target) {
				r[target] = buf
			}
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
