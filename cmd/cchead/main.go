package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	var filePath string
	var reader io.Reader
	maxLines := 10
	maxCharacters := 0
	useCharacters := false

	// Parse arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "-n") {
			// Handle -n flag
			if len(arg) > 2 {
				// Case: -n10 (value attached)
				numStr := arg[2:]
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatalf("Error converting string to int: %v", err)
				}
				maxLines = num
			} else {
				// Case: -n 10 (value as next argument)
				if i+1 < len(args) {
					i++ // Move to next argument
					num, err := strconv.Atoi(args[i])
					if err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					}
					maxLines = num
				} else {
					log.Fatalf("Flag -n requires a value")
				}
			}
		} else if strings.HasPrefix(arg, "-c") {
			// Handle -c flag
			useCharacters = true
			if len(arg) > 2 {
				// Case: -c100 (value attached)
				numStr := arg[2:]
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatalf("Error converting string to int: %v", err)
				}
				maxCharacters = num
			} else {
				// Case: -c 100 (value as next argument)
				if i+1 < len(args) {
					i++ // Move to next argument
					num, err := strconv.Atoi(args[i])
					if err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					}
					maxCharacters = num
				} else {
					log.Fatalf("Flag -c requires a value")
				}
			}
		} else if !strings.HasPrefix(arg, "-") {
			// This is the file path
			filePath = arg
		}
	}

	// Open file or use stdin
	if filePath == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	// Read and output based on flags
	if useCharacters {
		// Read specified number of characters
		buffer := make([]byte, maxCharacters)
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		fmt.Print(string(buffer[:n]))
	} else {
		// Read specified number of lines
		scanner := bufio.NewScanner(reader)
		lineCount := 0
		for scanner.Scan() && lineCount < maxLines {
			fmt.Println(scanner.Text())
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	}
}
