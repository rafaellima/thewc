package main

import (
	"fmt"
	"io"
	"log"
	"os"

	wc "github.com/rafaellima/thewc/pkg"
)

var cValue bool
var lValue bool
var wValue bool
var mValue bool
var hasFlag bool
var file string
var data string

func init() {
	for _, v := range os.Args {
		if v == "-c" {
			cValue = true
			hasFlag = true
		} else if v == "-l" {
			lValue = true
			hasFlag = true
		} else if v == "-w" {
			wValue = true
			hasFlag = true
		} else if v == "-m" {
			mValue = true
			hasFlag = true
		}
	}

	if len(os.Args) == 1 {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		data = string(content)
	} else if len(os.Args) == 2 && !hasFlag {
		file = os.Args[1]
		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		data = string(bytes)
	} else if len(os.Args) == 2 && hasFlag {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		data = string(content)
	} else if len(os.Args) > 2 {
		file = os.Args[2]
		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		data = string(bytes)
	} else {
		log.Fatal("missing arguments")
	}
}

func main() {
	counter := &wc.Wc{
		Data: data,
	}

	if cValue {
		fmt.Println(counter.CountChars())
	} else if lValue {
		fmt.Println(counter.CountLines())
	} else if wValue {
		fmt.Println(counter.CountWords())
	} else if mValue {
		fmt.Println(counter.CountBytes())
	} else {
		lineCount := counter.CountLines()
		wordCount := counter.CountWords()
		byteCount := counter.CountBytes()

		if file != "" {
			fmt.Println(lineCount, " ", wordCount, " ", byteCount, " ", file)
		} else {
			fmt.Println(lineCount, " ", wordCount, " ", byteCount)
		}
	}
}
