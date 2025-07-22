package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unix-tools/internal/cchead"
)

func main() {
	args := os.Args[1:]

	// Parse CLI args using helper so it's unit-testable.
	opts, err := cchead.ParseArgs(args)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var reader io.Reader
	if opts.FilePath == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(opts.FilePath)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()
		reader = file
	}

	var output string
	if opts.UseCharacters {
		output, err = cchead.HeadCharacters(reader, opts.MaxCharacters)
	} else {
		output, err = cchead.HeadLines(reader, opts.MaxLines)
	}
	if err != nil {
		log.Fatalf("Error reading: %v", err)
	}

	fmt.Print(output)
}
