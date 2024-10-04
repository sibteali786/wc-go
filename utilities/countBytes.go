package utilities

import (
	"io"
)

func CountBytes(r io.Reader) (int, error) {
	buffer := make([]byte, 1024) // 1 KB chunk
	byteCount := 0

	for {
		n, err := r.Read(buffer)
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
