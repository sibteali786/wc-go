package utilities

import "os"

func ReadFile(filename string) (*os.File, error) {
	rootPath, _ := os.Getwd()
	content, err := os.Open(rootPath + "/" + filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}
