package file

import (
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
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, PermDirectory); err != nil {
			return 0, err
		}
		return 0, err
	}

	f, err := os.Create(name)
	if err != nil {
		return 0, err
	}

	n, err := f.WriteString(s)
	if err != nil {
		return 0, err
	}

	if err = f.Chmod(PermFile); err != nil {
		return 0, err
	}

	return n, f.Close()
}
