package file

import (
	"bytes"
	"os"
	"testing"

	"text/template"
)

// Suite is a instance of Test struct
var Suite = &Test{Temp: "./tmp"}

// Test is a struct to help make tests
type Test struct {
	Temp string
}

// BuildPayloadFile is a func to help interpolate value in payload struct of file
func (tst *Test) BuildPayloadFile(data string, str interface{}, t *testing.T) string {
	bufferData := bytes.NewBuffer(nil)
	tmpl := template.Must(template.New("data").Parse(data))
	if err := tmpl.Execute(bufferData, str); err != nil {
		t.Fatal(err)
	}
	return bufferData.String()
}

// CreateTempFile is a func to create a file inside temp context
func (tst *Test) CreateTempFile(path, s string, t *testing.T) {
	if _, err := Build(path, s); err != nil {
		t.Fatal(err)
	}
}

// RemoveTemp is a func to remove temp folder
func (tst *Test) RemoveTemp(t *testing.T) {
	if err := os.RemoveAll(tst.Temp); err != nil {
		t.Fatal(err)
	}
}
