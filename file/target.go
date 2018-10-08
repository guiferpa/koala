package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type errAnyTargetFounded struct{}

func (e *errAnyTargetFounded) Error() string {
	return "any target founded"
}

// DefaultTag is a fallback value for tag input
const DefaultTag = "import"

// Target is a struct to save the target info
type Target struct {
	Tag     string
	Library string
}

func (t Target) String() string {
	return fmt.Sprintf("%s %s", t.Tag, t.Library)
}

// FindOutTargets is a func to find out target line in entry file
func FindOutTargets(tag, s string) ([]Target, error) {
	if tag == "" {
		tag = DefaultTag
	}
	targets := make([]Target, 0)
	scnr := bufio.NewScanner(bytes.NewBufferString(s))
	for scnr.Scan() {
		line := strings.TrimSpace(scnr.Text())
		if strings.HasPrefix(line, tag) {
			claimLibrary := strings.Split(line, " ")
			if len(claimLibrary) != 2 {
				continue
			}
			lineTag := claimLibrary[0]
			lineValue := claimLibrary[1]
			if lineTag == tag {
				targets = append(targets, Target{
					Tag:     lineTag,
					Library: lineValue,
				})
			}
		}
	}
	if len(targets) == 0 {
		return nil, &errAnyTargetFounded{}
	}
	return targets, nil
}

// ReplaceTargets is a func to replace targets by content from libraries
func ReplaceTargets(targets []Target, s string) (string, error) {
	for _, target := range targets {
		library, err := ioutil.ReadFile(target.Library)
		if err != nil {
			return "", err
		}
		s = strings.Replace(s, target.String(), string(library), -1)
	}
	return s, nil
}
