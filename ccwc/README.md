# ccwc - Character Count Word Count

ccwc is a simple command-line utility that counts the number of bytes in a specified text file. This project is inspired by the Unix `wc` command but focuses solely on byte counting.

## Project Structure

```
ccwc
├── cmd
│   └── ccwc
│       └── main.go
├── internal
│   └── ccwc
│       ├── ccwc.go
│       └── ccwc_test.go
├── testdata
│   └── test.txt
├── go.mod
└── README.md
```

## Installation

To install the ccwc utility, clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd ccwc
```

Then, build the application using the following command:

```bash
go build -o ccwc ./cmd/ccwc
```

## Usage

To use the ccwc utility, run the following command in your terminal:

```bash
./ccwc <path-to-file>
```

Replace `<path-to-file>` with the path to the text file you want to analyze. The application will output the number of bytes in the specified file.

## Testing

To run the unit tests for the ccwc functionality, use the following command:

```bash
go test ./internal/ccwc
```

This will execute the tests defined in `ccwc_test.go` to ensure the byte counting logic is functioning correctly.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.