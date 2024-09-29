package utilities

import (
	"bufio"
)

// using splitting technique

// func CountLines(filename string) (int, error) {
// 	content, err := ReadFile(filename)
// 	if err != nil {
// 		return 0, err
// 	}
// 	// convert bytes to string
// 	contentStr := string(content)
// 	// split by lines \n
// 	sliceContent := strings.Split(contentStr, "\n")
// 	// Return the number of lines
// 	return len(sliceContent) - 1, nil
// }

// since basis upon input stream, is more efficient in terms of large inputs as per memory

func CountLines(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Count lines
	linesCount := 0
	for scanner.Scan() {
		linesCount++
	}
	if err != nil {
		return 0, err
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return linesCount, nil
}
