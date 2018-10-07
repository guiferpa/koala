package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/guiferpa/koala/args"
	"github.com/guiferpa/koala/file"
)

var tag string

func init() {
	flag.StringVar(&tag, "tag", "import", "set a custom tag to mark a target, line that will be replaced by the assign content")
}

// ErrRequiredArg is a error type for input error
type errRequiredArg struct {
	arg    string
	reason error
}

func (err *errRequiredArg) Error() string {
	return fmt.Sprintf("%v argument is required, reason: %v", err.arg, err.reason)
}

func main() {
	flag.Parse()
	arguments := args.Parse(os.Args)
	tailargs := arguments.Tail()

	entry, err := arguments.First()
	if err != nil {
		log.Fatal(&errRequiredArg{arg: "entry", reason: err})
	}

	bundled, err := tailargs.First()
	if err != nil {
		log.Fatal(&errRequiredArg{arg: "bundled", reason: err})
	}

	payload, err := ioutil.ReadFile(path.Clean(entry))
	if err != nil {
		log.Fatal(err)
	}

	currentContent := string(payload)

	targets, err := file.FindOutTarget(tag, currentContent)
	if err != nil {
		log.Fatal(err)
	}

	newestContent, err := file.ReplaceTarget(targets, currentContent)
	if err != nil {
		log.Fatal(err)
	}

	n, err := file.Build(bundled, newestContent)
	if err != nil {
		log.Fatal(err)
	}

	templateSuccessMessage := "spelled successfully %v bytes at %v\n"
	absBundled, err := filepath.Abs(bundled)
	if err != nil {
		log.Printf(templateSuccessMessage, n, bundled)
	}
	log.Printf(templateSuccessMessage, n, absBundled)
}
