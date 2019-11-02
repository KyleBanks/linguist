package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	lang   string
	rounds int
)

func init() {
	flag.StringVar(&lang, "lang", "", "Two letter language code of the language to practise")
	flag.IntVar(&rounds, "rounds", 5, "The number of rounds to practise")
	flag.Parse()

	if lang == "" {
		fmt.Println("Missing required argument --lang")
		os.Exit(1)
	}
}

func main() {
	file, err := os.Open(fmt.Sprintf("./dataset/%s.txt", lang))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d, err := NewDataset(file)
	if err != nil {
		panic(err)
	}

	var count int
	buf := bufio.NewReader(os.Stdin)
	for _, native := range d.Native {
		fmt.Printf("Translate: %s\n", native)
		fmt.Print("> ")

		input, err := buf.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		var correct bool
		translations := d.Translations[native]
		for _, trans := range translations {
			correct = IsCorrect(string(input), trans)
			if correct {
				break
			}
		}

		if correct {
			fmt.Println("✓")
		} else {
			fmt.Printf("✘ %s\n", strings.Join(translations, "\n  "))
		}
		fmt.Println("")

		count++
		if count == rounds {
			break
		}
	}
}
