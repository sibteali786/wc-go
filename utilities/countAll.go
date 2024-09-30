package utilities

func CountAll(filename string) ([3]int, error) {
	// CountBytes, CountLines, and CountWords should be implemented elsewhere.
	bytes, errBytes := CountBytes(filename)
	if errBytes != nil {
		return [3]int{}, errBytes // return zero value for array and error
	}

	lines, errLines := CountLines(filename)
	if errLines != nil {
		return [3]int{}, errLines // return zero value for array and error
	}

	words, errWords := CountWords(filename)
	if errWords != nil {
		return [3]int{}, errWords // return zero value for array and error
	}

	// No errors, construct the array and return
	arr := [3]int{lines, words, bytes}
	return arr, nil
}
