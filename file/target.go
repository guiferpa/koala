package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Target struct {
	Tag     string
	Library string
}

func (t Target) String() string {
	return fmt.Sprintf("%s %s", t.Tag, t.Library)
}

func FindOutTarget(tag, s string) ([]Target, error) {
	targets := make([]Target, 0)
	scnr := bufio.NewScanner(bytes.NewBufferString(s))
	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), tag) {
			claimLibrary := strings.Split(scnr.Text(), " ")
			if len(claimLibrary) > 2 {
				return nil, errors.New("import syntax wrong")
			}
			if len(claimLibrary) == 2 {
				targets = append(targets, Target{
					Tag:     claimLibrary[0],
					Library: claimLibrary[1],
				})
			}
		}
	}
	return targets, nil
}

func ReplaceTarget(targets []Target, s string) (string, error) {
	for _, target := range targets {
		library, err := ioutil.ReadFile(target.Library)
		if err != nil {
			return "", err
		}
		s = strings.Replace(s, target.String(), string(library), -1)
	}
	return s, nil
}
