package file

import (
	"fmt"
	"testing"
)

func TestFindOutTargets(t *testing.T) {
	tag := "test"
	data := []byte(`
test 1
test 2
 test 3
	test 4
t 4
te 5
tes 6
test 7
tes 8
te 9
t 10
testing 11
testin 12
testi 13
`)
	result, err := FindOutTargets(tag, string(data))
	if err != nil {
		t.Error(err)
		return
	}
	expected := []Target{
		{Tag: "test", Library: "1"},
		{Tag: "test", Library: "2"},
		{Tag: "test", Library: "7"},
	}
	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

type tmpLibrary struct {
	Path  string
	Value string
}

func TestReplaceTags(t *testing.T) {
	lib := tmpLibrary{
		Path:  fmt.Sprintf("%v/lib/1", Suite.Temp),
		Value: "hello 1",
	}

	Suite.CreateTempFile(lib.Path, lib.Value, t)
	defer Suite.RemoveTemp(t)

	in := []Target{{Tag: "test", Library: lib.Path}}
	data := `
test {{.Path}}

hello test
`
	result, err := ReplaceTargets(in, Suite.BuildPayloadFile(data, lib, t))
	if err != nil {
		t.Error(err)
		return
	}

	expectedData := `
{{.Value}}

hello test
`
	expected := Suite.BuildPayloadFile(expectedData, lib, t)
	if result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}
