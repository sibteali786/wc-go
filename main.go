package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"wcgo/utilities"
)

func main() {
	// Defining a boolean flag -l to count lines
	lines := flag.Bool("l", false, "Count Lines")
	bytes := flag.Bool("c", false, "Count Bytes")
	words := flag.Bool("w", false, "Count Words")
	characters := flag.Bool("m", false, "Count Characters")
	// parsing the flags provided by user
	flag.Parse()
	fmt.Println(Count(os.Stdin, *lines, *bytes, *words, *characters))
	// fmt.Println(Count(os.Stdin, *lines))
}
func Count(r io.Reader, lines bool, bytes bool, words bool, characters bool) (int, error) {
	if lines {
		result, err := utilities.CountLines("test.txt")
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if words {
		result, err := utilities.CountWords("test.txt")
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if bytes {
		result, err := utilities.CountBytes("test.txt")
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if characters {
		result, err := utilities.CountCharacters("test.txt")
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	return 0, nil
}
