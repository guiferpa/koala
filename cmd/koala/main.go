package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/guiferpa/koala/file"
)

var (
	entry   string
	bundled string
	tag     string
)

func init() {
	flag.StringVar(&entry, "entry", "./entry", "main file to reading")
	flag.StringVar(&bundled, "bundled", "./bin/bundled", "output file for export bundle")
	flag.StringVar(&tag, "tag", "import", "set a custom tag to mark a target, line that will be replaced by the assign content")
}

func main() {
	flag.Parse()

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
	fmt.Println(n)
}
