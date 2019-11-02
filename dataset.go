package main

import (
	"bufio"
	"fmt"
	"go/build"
	"io"
	"os"
	"strings"
)

type Dataset struct {
	Native       []string
	Translations map[string][]string
}

func NewDataset(r io.Reader) (*Dataset, error) {
	d := Dataset{
		Translations: make(map[string][]string),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		native := Sanitize(parts[0])
		trans := Sanitize(parts[1])

		if _, ok := d.Translations[native]; !ok {
			d.Native = append(d.Native, native)
			d.Translations[native] = make([]string, 0)
		}
		d.Translations[native] = append(d.Translations[native], trans)
	}

	return &d, scanner.Err()
}

func getLangFile(lang string) (*os.File, error) {
	// TODO better means of locating dataset files
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	return os.Open(fmt.Sprintf("%s/src/github.com/KyleBanks/linguist/dataset/%s.txt", gopath, lang))
}
