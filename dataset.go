package main

import (
	"bufio"
	"io"
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
