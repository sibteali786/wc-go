package utilities

import (
	"bufio"
	"io"
)

func CountCharacters(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	charCount := 0
	for scanner.Scan() {
		charCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return charCount, nil
}
