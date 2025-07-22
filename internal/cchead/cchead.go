package cchead

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// Options represents the parameters for the head operation.
// If UseCharacters is true, MaxCharacters is respected and MaxLines is ignored.
// If UseCharacters is false, MaxLines is respected.
// Either field can be 0 to indicate the default (10 lines or N characters).
// FilePath holds an optional path to a file to read; if empty, the caller should
// fall back to STDIN.
// This mirrors the behaviour of GNU head.
type Options struct {
	MaxLines      int
	MaxCharacters int
	UseCharacters bool
	FilePath      string
}

// DefaultMaxLines provides the default number of lines when -n is not supplied.
const DefaultMaxLines = 10

// ParseArgs converts the CLI style arguments (excluding program name) into an
// Options struct. It recognises the following patterns:
//
//	-n <N>   or -n<N>
//	-c <N>   or -c<N>
//
// Any non-flag argument is treated as the file path. Multiple file paths are
// not supported and will return an error.
func ParseArgs(args []string) (Options, error) {
	opts := Options{MaxLines: DefaultMaxLines}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-n") {
			// Handle -n flag
			if strings.HasPrefix(arg, "-n") && len(arg) > 2 { // -n10
				num, err := strconv.Atoi(arg[2:])
				if err != nil {
					return opts, err
				}
				opts.MaxLines = num
			} else {
				// -n 10
				if i+1 >= len(args) {
					return opts, errors.New("flag -n requires a value")
				}
				i++
				num, err := strconv.Atoi(args[i])
				if err != nil {
					return opts, err
				}
				opts.MaxLines = num
			}
		} else if strings.HasPrefix(arg, "-c") {
			// Handle -c flag
			opts.UseCharacters = true
			if len(arg) > 2 { // -c100
				num, err := strconv.Atoi(arg[2:])
				if err != nil {
					return opts, err
				}
				opts.MaxCharacters = num
			} else {
				if i+1 >= len(args) {
					return opts, errors.New("flag -c requires a value")
				}
				i++
				num, err := strconv.Atoi(args[i])
				if err != nil {
					return opts, err
				}
				opts.MaxCharacters = num
			}
		} else if strings.HasPrefix(arg, "-") {
			return opts, errors.New("unknown flag: " + arg)
		} else {
			// File path
			if opts.FilePath != "" {
				return opts, errors.New("multiple file paths provided")
			}
			opts.FilePath = arg
		}
	}

	// For -c ensure we have a positive char count
	if opts.UseCharacters && opts.MaxCharacters <= 0 {
		return opts, errors.New("character count must be > 0")
	}

	// For lines ensure positive
	if !opts.UseCharacters && opts.MaxLines <= 0 {
		return opts, errors.New("line count must be > 0")
	}

	return opts, nil
}

// HeadLines reads up to maxLines from r and returns them as a single string.
func HeadLines(r io.Reader, maxLines int) (string, error) {
	scanner := bufio.NewScanner(r)
	var builder strings.Builder
	lineCount := 0
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteByte('\n')
		lineCount++
		if lineCount >= maxLines {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return builder.String(), nil
}

// HeadCharacters reads up to maxCharacters bytes from r and returns them as a
// string.
func HeadCharacters(r io.Reader, maxCharacters int) (string, error) {
	buf := make([]byte, maxCharacters)
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf[:n]), nil
}
