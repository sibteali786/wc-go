package utilities

import (
	"bufio"
	"os"
)

func HandleStdIn() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}
