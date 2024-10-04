package main

import (
	"os"
	"testing"
	"wcgo/utilities"
)

func TestCountBytes(t *testing.T) {
	expected := 342190 // bytes of the test.txt file
	result, err := utilities.CountBytes("test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Expected %d , got %d", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	expected := 7145
	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	result, err := utilities.CountLines(file)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCountWords(t *testing.T) {
	expected := 58164
	result, err := utilities.CountWords("test.txt")
	if err != nil {
		t.Error(err)
		return
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCountCharacters(t *testing.T) {
	expected := 339292
	result, err := utilities.CountCharacters("test.txt")
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
func TestCountAll(t *testing.T) {
	expectedLines := 7145
	expectedWords := 58164
	expectedBytes := 342190
	arr, err := utilities.CountAll("test.txt")
	if err != nil {
		t.Error(err)
	}
	for i, values := range arr {
		if i == 0 && values != expectedLines {
			t.Errorf("Expected %d, got %d", expectedLines, values)
		}
		if i == 1 && values != expectedWords {
			t.Errorf("Expected %d, got %d", expectedWords, values)
		}
		if i == 2 && values != expectedBytes {
			t.Errorf("Expected %d, got %d", expectedBytes, values)
		}
	}
}
