package utilities

func CountBytes(filename string) (int, error) {
	content, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(content), nil
}
