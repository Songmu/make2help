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
		"task1": []string{"task1 desu"},
		"task2": []string{},
		"task3": []string{"task3 desu", "multi line"},
	}

	if !reflect.DeepEqual(r, expect) {
		t.Errorf("somthing went wrong: %#v", r)
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
