package make2help

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	r, err := scan("testdata/Makefile")

	if err != nil {
		t.Errorf("err should nil but: %s", err)
	}

	expect := rules{
		"task1":  []string{"task1 desu"},
		"task2":  []string{},
		"task3":  []string{"task3 desu", "multi line"},
		"task4":  []string{"task4 desuyo"},
		"task5":  []string{"task5 no phony"},
		"task6":  []string{"task6 suffix whitespace"},
		"task7%": []string{"task7 pattern rule"},
	}

	if !reflect.DeepEqual(r, expect) {
		t.Errorf("somthing went wrong\n   got: %#v\nexpect: %#v", r, expect)
	}
}

func TestMerge(t *testing.T) {
	r1 := rules{
		"task1": []string{"task1 desu"},
		"task2": []string{},
		"task3": []string{"task3 desu", "multi line"},
	}

	r2 := rules{
		"task2": []string{"task2 desu"},
		"task4": []string{},
	}

	r3 := r1.merge(r2)

	expect := rules{
		"task1": []string{"task1 desu"},
		"task2": []string{"task2 desu"},
		"task3": []string{"task3 desu", "multi line"},
		"task4": []string{},
	}

	if !reflect.DeepEqual(r3, expect) {
		t.Errorf("somthing went wrong: %#v", r3)
	}
}
