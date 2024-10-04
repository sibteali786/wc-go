package utilities

import (
	"bufio"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	// Count lines
	linesCount := 0
	for scanner.Scan() {
		linesCount++
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return linesCount, nil
}
