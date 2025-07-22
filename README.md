## Implementing UNIX commands from scratch

### ✅Completed

- wc: word count

### Todo

- cat
- cut

---

## Getting started

These tools are ordinary Go programs. You can run them right away (without installing anything system-wide) as long as you have Go ≥ 1.23 on your PATH.

### Quick run – no compile step

```
# Show first 10 lines of a file (default)
go run ./cmd/cchead           testdata/test.txt

# First 5 lines
go run ./cmd/cchead -n 5      testdata/test.txt

# First 100 characters coming from STDIN
echo "hello world" | go run ./cmd/cchead -c 100

# Count bytes, words and lines in a file (default flags -c -w -l)
go run ./cmd/ccwc             testdata/test.txt

# Only word count
go run ./cmd/ccwc -w          testdata/test.txt
```

### Building standalone binaries

```
go build -o cchead ./cmd/cchead
go build -o ccwc   ./cmd/ccwc

./cchead -n 20 somefile.txt
./ccwc   -l somefile.txt
```

### Running the unit-tests

```
go test ./...
```

---

Each command has two parts:

1. **`cmd/<tool>/main.go`** – tiny CLI wrapper (flag parsing, opening files / STDIN).
2. **`internal/<tool>`** – pure functions that implement the real logic. They’re easy to unit-test and cannot be imported from outside the module (a Go convention).

Feel free to open an issue or PR if you spot a bug or would like to add a new UNIX utility.
