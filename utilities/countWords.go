package utilities

import (
	"bufio"
	"io"
)

func CountWords(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wordsCount := 0
	for scanner.Scan() {
		wordsCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return wordsCount, nil
}
