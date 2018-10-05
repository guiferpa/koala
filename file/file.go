package file

import (
	"os"
	"path"
)

const (
	PermDirectory = 0777
	PermFile      = 0644
)

func Build(name, s string) error {
	dir := path.Dir(name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, PermDirectory); err != nil {
			return err
		}
		return err
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(s)
	file.Chmod(PermFile)

	return nil
}
