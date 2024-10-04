package utilities

import (
	"bufio"
	"io"
	"unicode"
)

func CountAll(r io.Reader) ([3]int, error) {
	var linesCount, wordsCount, byteCount int

	reader := bufio.NewReader(r)
	inWord := false

	for {
		runeValue, size, err := reader.ReadRune()

		// Handle errors
		if err != nil {
			if err == io.EOF {
				if inWord {
					wordsCount++
				}
				break
			} else {
				return [3]int{}, err
			}
		}

		// Increment byte count by the size of the rune
		byteCount += size

		// Check if the rune is a line break
		if runeValue == '\n' {
			linesCount++
		}

		// Handle word counting using unicode.IsSpace to determine boundaries
		if unicode.IsSpace(runeValue) {
			if inWord {
				wordsCount++
			}
			inWord = false
		} else {
			inWord = true
		}
	}

	// Return counts for lines, words, bytes, and characters
	return [3]int{linesCount, wordsCount, byteCount}, nil
}
