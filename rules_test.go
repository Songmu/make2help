package make2help

import "testing"

func TestString(t *testing.T) {

	r := rules{
		"task1": []string{"task1 desu"},
		"task2": []string{},
		"task3": []string{"task3 desu", "multi line"},
	}

	result := r.string(false, false)
	expect := `task1:             task1 desu
task3:             task3 desu
                   multi line
`

	if result != expect {
		t.Errorf("result is not expected one: %s", result)
	}
}
