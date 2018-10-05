package file

import (
	"os"
	"path"
	"testing"
)

const base = "./tmp"

func TestBuildFileIfDirNotExists(t *testing.T) {
	name := path.Join(base, "1", "llama.txt")
	txt := "testing..."
	if err := Build(name, txt); err != nil {
		t.Error(err)
	}
	defer func(d string) {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
	}(base)
}

func TestBuildFileIfDirAlreadyExists(t *testing.T) {
	dir := path.Join(base, "2")
	name := path.Join(dir, "koala.txt")
	txt := "testing..."
	if err := os.MkdirAll(dir, PermDirectory); err != nil {
		t.Fatal(err)
	}
	if err := Build(name, txt); err != nil {
		t.Error(err)
	}
	defer func(d string) {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
	}(base)
}
