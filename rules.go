package make2help

import (
	"bytes"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
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
	indent, tasks := r.indentAndRules(all)
	colorfunc := ansi.ColorFunc("cyan")
	for _, task := range tasks {
		helplines := r[task]
		msg := task
		if colorful {
			msg = colorfunc(msg)
		}
		buf.WriteString(msg)
		buf.WriteString(":")
		restIndent := indent - len(task) - 1
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
			// delimiter(Kolon) and space = 2 width
			if len(k)+2 > indent {
				indent = len(k) + 2
			}
			tasks = append(tasks, k)
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return strings.ToLower(tasks[i]) < strings.ToLower(tasks[j])
	})
	return indent, tasks
}
