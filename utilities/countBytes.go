package utilities

import (
	"io"
)

func CountBytes(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read the file in chunks
	buffer := make([]byte, 1024) // Read 1 KB at a time
	byteCount := 0

	for {
		n, err := file.Read(buffer) // Read reads up to len(b) bytes from the File and stores them in b. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF
		if err != nil && err != io.EOF {
			return 0, err
		}
		byteCount += n
		if err == io.EOF {
			break
		}
	}

	return byteCount, nil
}
