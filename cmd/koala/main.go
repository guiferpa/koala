package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/guiferpa/koala/args"
	"github.com/guiferpa/koala/file"
)

// ErrRequiredArg is a error type for input error
type errRequiredArg struct {
	arg    string
	reason error
}

func (err *errRequiredArg) Error() string {
	return fmt.Sprintf("%v argument is required, reason: %v", err.arg, err.reason)
}

func main() {
	entryargs := args.Parse(os.Args)
	bundledargs := entryargs.Tail()
	flagargs := bundledargs.Tail()

	entry, err := entryargs.First()
	if err != nil {
		log.Fatal(&errRequiredArg{arg: "entry", reason: err})
	}

	bundled, err := bundledargs.First()
	if err != nil {
		log.Fatal(&errRequiredArg{arg: "bundled", reason: err})
	}

	tag, err := flagargs.First()
	if err != nil {
		log.Println(fmt.Sprintf("no custom tag then by default will be use <%v> as tag", file.DefaultTag))
	}

	payload, err := ioutil.ReadFile(path.Clean(entry))
	if err != nil {
		log.Fatal(err)
	}

	currentContent := string(payload)

	targets, err := file.FindOutTargets(tag, currentContent)
	if err != nil {
		log.Fatal(err)
	}

	newestContent, err := file.ReplaceTargets(targets, currentContent)
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
