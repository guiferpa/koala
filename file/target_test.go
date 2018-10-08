package file

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindOutTargets(t *testing.T) {
	tag := "join"
	data := `
join 1
join 2
 join 3
	join 4
j 5
jo 6
joi 7
join ('8')
joi 9
jo 10
j 11
joi 12
join('13')

testing to pkg=file

joining to pkg=file

import test = "testing.." (
	variable t = 1000
	join test
)
`
	result, err := FindOutTargets(tag, data)
	if err != nil {
		t.Error(err)
		return
	}
	expected := []Target{
		{Tag: "join", Library: "1"},
		{Tag: "join", Library: "2"},
		{Tag: "join", Library: "3"},
		{Tag: "join", Library: "4"},
		{Tag: "join", Library: "('8')"},
		{Tag: "join", Library: "test"},
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

func TestFindOutTargetsWithErrAnyTargetFounded(t *testing.T) {
	cases := []struct {
		Tag  string
		Data string
	}{
		{
			Tag: "include",
			Data: `
import { test } from './suite';

Test.do(func(tester=test) {
	tester lib/fakerfunc
})
`},
		{
			Tag: "include",
			Data: `
include('test');
include(' test ');
`},
	}

	for _, test := range cases {
		_, err := FindOutTargets(test.Tag, test.Data)
		if _, ok := err.(*errAnyTargetFounded); !ok {
			t.Error("unexpected error type", reflect.TypeOf(err))
		}
	}
}

func TestErrAnyTargetFounded(t *testing.T) {
	expected := "any target founded"
	err := &errAnyTargetFounded{}
	if result := err.Error(); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}
