package main

import (
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	content, err := os.ReadFile(rootPath + "/test.txt")
	fmt.Println(content, err)
}
