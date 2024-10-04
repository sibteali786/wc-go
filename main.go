package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"wcgo/utilities"
)

func main() {
	// Defining the flags
	lines := flag.Bool("l", false, "Count Lines")
	bytes := flag.Bool("c", false, "Count Bytes")
	words := flag.Bool("w", false, "Count Words")
	characters := flag.Bool("m", false, "Count Characters")
	filePath := flag.String("f", "", "File Path")
	flag.Parse()

	// If no specific flag is set, count all metrics

	// If filePath is empty, use Stdin
	var reader io.Reader
	if *filePath == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		reader = file
	}
	if !*lines && !*bytes && !*words && !*characters {
		result, err := utilities.CountAll(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		strSlice := make([]string, len(result))
		for i, v := range result {
			strSlice[i] = strconv.Itoa(v)
		}
		if *filePath != "" {
			strSlice = append(strSlice, *filePath)
		}
		fmt.Println(strings.Join(strSlice, " "))
	} else {
		result, err := Count(reader, *lines, *bytes, *words, *characters)
		if err != nil {
			fmt.Println(err)
			return
		}

		strResult := strconv.Itoa(result) + " " + *filePath
		fmt.Println(strResult)
	}

}

func Count(r io.Reader, lines bool, bytes bool, words bool, characters bool) (int, error) {
	if lines {
		result, err := utilities.CountLines(r)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if words {
		result, err := utilities.CountWords(r)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if bytes {
		result, err := utilities.CountBytes(r)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	if characters {
		result, err := utilities.CountCharacters(r)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	return 0, nil
}
