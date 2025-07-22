package cchead

import (
	"os"
	"testing"
)

func TestHeadCommand(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"../../testdata/test_small.txt", "This file is unintentionally left blank.\n"},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			file, err := os.Open(test.filename)
			if err != nil {
				t.Fatalf("Failed to open file %q: %v", test.filename, err)
			}
			defer file.Close()

			got, err := HeadLines(file, 5)
			if err != nil {
				t.Fatalf("HeadLines(%q): returned an error: %v", test.filename, err)
			}
			if got != test.expected {
				t.Errorf("HeadLines(%q) = %q; want %q", test.filename, got, test.expected)
			}
		})
	}
}
