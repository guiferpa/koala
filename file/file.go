package file

import (
	"io/ioutil"
	"os"
	"path"
)

const (
	// PermDirectory is a alias for permission of bundle directory
	PermDirectory = 0777
	// PermFile is a alias for permission of bundle file
	PermFile = 0755
)

// Build is the func to create the bundle file
func Build(name, s string) (int, error) {
	dir := path.Dir(name)
	if err := os.MkdirAll(dir, PermDirectory); err != nil {
		return 0, err
	}

	if err := ioutil.WriteFile(name, []byte(s), PermFile); err != nil {
		return 0, err
	}

	statFile, err := os.Stat(name)
	if err != nil {
		return 0, err
	}

	return int(statFile.Size()), nil
}
