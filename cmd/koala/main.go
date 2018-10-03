package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	entrypath string
	entryfile string
	outpath   string
	outfile   string
	target    string
)

func init() {
	flag.StringVar(&entrypath, "entrypath", ".", "set a custom entrypath for your entryfile")
	flag.StringVar(&entryfile, "entryfile", "main", "entryfile is where check lib imports")
	flag.StringVar(&outpath, "outpath", "./bin", "set a custom outpath for your distro")
	flag.StringVar(&outfile, "outfile", "main", "outfile for export bundle")
	flag.StringVar(&target, "target", "import", "set a custom target to mark line that replace content")
}

func main() {
	flag.Parse()

	entrypoint := fmt.Sprintf("%s/%s", entrypath, entryfile)
	output := fmt.Sprintf("%s/%s", outpath, outfile)

	fileContent, err := ioutil.ReadFile(entrypoint)
	if err != nil {
		log.Fatalln(err)
	}

	originFile := bytes.NewBuffer(fileContent)
	copyFile := bytes.NewBufferString(originFile.String())

	targetLines, err := getTargetLines(target, copyFile)
	if err != nil {
		panic(err)
	}

	dstFile, err := replaceByTargetLines(targetLines, originFile.String())
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		os.Mkdir(outpath, 0777)
	}

	if err = ioutil.WriteFile(output, dstFile, 0644); err != nil {
		panic(err)
	}
}

func getTargetLines(target string, file *bytes.Buffer) ([]string, error) {
	if target == "" {
		return nil, errors.New("the target is empty")
	}
	targetLines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), target) {
			claimLibrary := strings.Split(scanner.Text(), " ")
			if len(claimLibrary) != 2 {
				return nil, errors.New("import syntax wrong")
			}
			targetLines = append(targetLines, scanner.Text())
		}
	}
	return targetLines, scanner.Err()
}

func replaceByTargetLines(targetLines []string, originFile string) ([]byte, error) {
	for _, targetLine := range targetLines {
		splittedTargetLine := strings.Split(targetLine, " ")
		contentLib, err := ioutil.ReadFile(splittedTargetLine[1])
		if err != nil {
			return nil, err
		}
		originFile = strings.Replace(originFile, targetLine, string(contentLib), -1)
	}
	return []byte(originFile), nil
}

func replaceFileFromImportToSource(fileSource string, libraries []string) (string, error) {
	for _, lib := range libraries {
		splittedLibraryImport := strings.Split(lib, " ")
		libSource, err := ioutil.ReadFile(splittedLibraryImport[1])
		if err != nil {
			return "", err
		}
		fileSource = strings.Replace(fileSource, lib, string(libSource), -1)
	}
	return fileSource, nil
}
