package file

import (
	"os"
	"path"
	"testing"
)

func TestBuildFileIfDirNotExists(t *testing.T) {
	name := path.Join(Suite.Temp, "1", "llama.txt")
	txt := "testing..."

	result, err := Build(name, txt)
	if err != nil {
		t.Error(err)
	}

	expected := len([]byte(txt))
	if result != expected {
		t.Errorf("unexpected result, result: %v bytes, expected: %v bytes", result, expected)
	}

	defer Suite.RemoveTemp(t)
}

func TestBuildFileIfDirAlreadyExists(t *testing.T) {
	dir := path.Join(Suite.Temp, "2")
	name := path.Join(dir, "koala.txt")
	txt := "testing..."

	if err := os.MkdirAll(dir, PermDirectory); err != nil {
		t.Fatal(err)
	}

	result, err := Build(name, txt)
	if err != nil {
		t.Error(err)
	}

	expected := len([]byte(txt))
	if result != expected {
		t.Errorf("unexpected result, result: %v bytes, expected: %v bytes", result, expected)
	}

	defer Suite.RemoveTemp(t)
}
