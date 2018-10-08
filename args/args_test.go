package args

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	in := []string{"a", "-b=test", "-c", "--d=other-test", "e", "f"}
	expected := "a"
	args := Parse(in)
	if result := args.Binary(); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestBinaryFromTail(t *testing.T) {
	in := []string{"binary", "a", "b", "c", "d"}
	expected := "binary"
	args := Parse(in).Tail()
	if result := args.Binary(); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestBinaryFromTailWithOnlyOneArg(t *testing.T) {
	in := []string{"binary", "a"}
	expected := "binary"
	args := Parse(in).Tail()
	if result := args.Binary(); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestCorretArgumentsPassed(t *testing.T) {
	in := []string{"binary", "a", "b", "c", "d"}
	expected := fmt.Sprintf("%v", []string{"a", "b", "c", "d"})
	args := Parse(in)
	if result := fmt.Sprintf("%v", args.args); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestFirstArguments(t *testing.T) {
	in := []string{"binary", "a", "b", "c", "d"}
	expected := "a"
	args := Parse(in)
	result, err := args.First()
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestTailArguments(t *testing.T) {
	in := []string{"binary", "a", "b", "c", "d"}
	expected := fmt.Sprintf("%v", []string{"b", "c", "d"})
	args := Parse(in).Tail()
	if result := fmt.Sprintf("%v", args.args); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestTailWithOnlyOneArg(t *testing.T) {
	in := []string{"binary", "a"}
	expected := fmt.Sprintf("%v", []string{})
	args := Parse(in).Tail()
	if result := fmt.Sprintf("%v", args.args); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestFirstWithOnlyBinary(t *testing.T) {
	in := []string{"binary"}
	_, err := Parse(in).First()
	if err == nil {
		t.Error("unexpected behavior, error expected")
		return
	}
	if _, ok := err.(*errEndOfArgs); !ok {
		t.Error("unexpected error type")
	}
}

func TestFistArgumentFromTail(t *testing.T) {
	in := []string{"binary", "a", "b", "c", "d"}
	expected := "b"
	args := Parse(in).Tail()
	result, err := args.First()
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestSanitizeFlags(t *testing.T) {
	in := []string{"binary", "a", "-b=test", "c", "--d=other-test"}
	expected := fmt.Sprintf("%v", []string{"a", "c"})
	args := Parse(in)
	if result := fmt.Sprintf("%v", args.args); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestSanitizeFlagsFromTail(t *testing.T) {
	in := []string{"a", "-b=test", "-c", "--d=other-test", "e", "f"}
	expected := fmt.Sprintf("%v", []string{"f"})
	args := Parse(in).Tail()
	if result := fmt.Sprintf("%v", args.args); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}

func TestErrEndOfArgs(t *testing.T) {
	expected := "end of arguments"
	err := &errEndOfArgs{}
	if result := err.Error(); result != expected {
		t.Errorf("unexpected result, result: %v, expected: %v", result, expected)
	}
}
