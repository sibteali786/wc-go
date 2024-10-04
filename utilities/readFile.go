package utilities

import (
	"os"
	"testing"
)

func ReadFile(t *testing.T, filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	return file
}
