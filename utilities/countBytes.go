package utilities

import (
	"os"
)

func CountBytes(filename string) (int, error) {
	rootPath, _ := os.Getwd()
	content, err := os.ReadFile(rootPath + "/" + filename)
	if err != nil {
		return 0, err
	}
	return len(content), nil
}
