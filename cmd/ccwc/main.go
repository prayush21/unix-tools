package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unix-tools/internal/ccwc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ccwc [-c|-w|-l|-m] [file]")
		os.Exit(1)
	}

	args := os.Args[1:]
	flags := []string{}
	var filePath string
	readFromStdin := false

	// Separate flags and file path
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			filePath = arg
		}
	}

	if len(flags) == 0 {
		flags = []string{"-c", "-w", "-l"}
	}

	var reader io.Reader
	if filePath == "" {
		// No file specified, read from stdin
		readFromStdin = true
		reader = os.Stdin
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	for _, flag := range flags {
		switch flag {
		case "-c":
			byteCount, err := ccwc.CountBytesReader(reader)
			if err != nil {
				fmt.Printf("Error reading: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Byte count: %d\n", byteCount)
		case "-l":
			lineCount, err := ccwc.CountLinesReader(reader)
			if err != nil {
				fmt.Printf("Error reading: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Line count: %d\n", lineCount)
		case "-w":
			wordCount, err := ccwc.CountWordsReader(reader)
			if err != nil {
				fmt.Printf("Error reading: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Word count: %d\n", wordCount)
		case "-m":
			charCount, err := ccwc.CountRunesReader(reader)
			if err != nil {
				fmt.Printf("Error reading: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Char count: %d\n", charCount)
		}
		// Reset reader if reading from stdin is not possible, otherwise re-open file for each flag
		if !readFromStdin && filePath != "" {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			reader = file
		}
	}
}
