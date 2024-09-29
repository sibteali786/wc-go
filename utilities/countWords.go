package utilities

import (
	"bufio"
)

func CountWords(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Use a split function to split by words
	scanner.Split(bufio.ScanWords)
	wordsCount := 0
	for scanner.Scan() {
		wordsCount++
	}
	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return wordsCount, nil
}
