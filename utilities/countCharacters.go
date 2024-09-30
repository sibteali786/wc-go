package utilities

import "bufio"

func CountCharacters(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
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
