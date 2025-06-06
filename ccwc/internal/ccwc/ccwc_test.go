package ccwc

import (
	// "os"
	"testing"
)

func TestCountBytes(t *testing.T) {
	tests := []struct {
		filename string
		expected int64
	}{
		{"./testdata/test_small.txt", 40}, // Adjust the expected value based on the content of test_small.txt
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			got, err := CountBytes(test.filename)
			if err != nil {
				t.Fatalf("CountBytes(%q) returned an error: %v", test.filename, err)
			}
			if got != test.expected {
				t.Errorf("CountBytes(%q) = %d; want %d", test.filename, got, test.expected)
			}
		})
	}
}
