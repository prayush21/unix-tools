package ccwc

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// CountBytes takes a file path and returns the number of bytes in that file.
func CountBytes(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return int64(len(content)), nil
}

// CountLines takes a file path and returns the number of lines in that file.
func CountLines(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return -1, err
	}
	lines := strings.Split(string(content), "\n")

	return int64(len(lines)), nil

}

// CountWords takes a file path and returns the number of words in that file.
func CountWords(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return -1, err
	}
	// words := strings.Split(string(content), " ")
	words := strings.Fields(string(content))
	// fmt.Println("Maybe Characters/Runes: ", len(strings.ToValidUTF8(string(content))))
	return int64(len(words)), nil
}

// CountRunes takes a file path and returns the number of runes (Unicode characters) in that file.
func CountRunes(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return -1, err
	}
	// words := strings.Split(string(content), " ")
	chars := []rune(string(content))
	// fmt.Println("Maybe Characters/Runes: ", len(strings.ToValidUTF8(string(content))))
	return int64(len(chars)), nil
}

// CountBytesReader returns the number of bytes in the reader.
func CountBytesReader(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

// CountLinesReader returns the number of lines in the reader.
func CountLinesReader(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

// CountWordsReader returns the number of words in the reader.
func CountWordsReader(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

// CountRunesReader returns the number of runes (Unicode characters) in the reader.
func CountRunesReader(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
