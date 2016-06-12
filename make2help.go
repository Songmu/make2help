package make2help

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var ruleReg = regexp.MustCompile(`^(\S+):`)

func scan(filepath string) (rules, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
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
			r[matches[1]] = buf
		}
		if len(buf) > 0 {
			buf = []string{}
		}
	}
	if err = sc.Err(); err != nil {
		return nil, errors.Wrap(err, "scan failed")
	}
	return r, nil
}
