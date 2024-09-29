package utilities

import "os"

func ReadFile(filename string) ([]byte, error) {
	rootPath, _ := os.Getwd()
	content, err := os.ReadFile(rootPath + "/" + filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}
