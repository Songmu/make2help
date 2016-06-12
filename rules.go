package make2help

import (
	"bytes"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/patrickmn/sortutil"
)

type rules map[string][]string

func (r rules) merge(r2 rules) rules {
	for k, v := range r2 {
		r[k] = v
	}
	return r
}

func (r rules) string(all, colorful bool) string {
	var buf bytes.Buffer
	header := "Available rules:"
	if colorful {
		header = "\033[1m" + header + "\033[0m"
	}
	buf.WriteString(header)
	buf.WriteString("\n\n")

	indent, tasks := r.indentAndRules(all)
	colorfunc := ansi.ColorFunc("cyan")
	for _, task := range tasks {
		helplines := r[task]
		msg := task
		if colorful {
			msg = colorfunc(msg)
		}
		buf.WriteString(msg)
		restIndent := indent - len(task)
		if len(helplines) < 1 {
			buf.WriteString("\n")
		}
		for _, helpline := range helplines {
			buf.WriteString(strings.Repeat(" ", restIndent))
			buf.WriteString(helpline)
			buf.WriteString("\n")
			restIndent = indent
		}
	}
	return buf.String()
}

func (r rules) indentAndRules(all bool) (int, []string) {
	indent := 19
	var tasks []string
	for k, v := range r {
		if all || len(v) > 0 {
			if len(k)+1 > indent {
				indent = len(k) + 1
			}
			tasks = append(tasks, k)
		}
	}
	sortutil.CiAsc(tasks)
	return indent, tasks
}
