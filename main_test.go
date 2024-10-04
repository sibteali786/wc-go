package main

import (
	"testing"
	"wcgo/utilities"
)

func TestCountLines(t *testing.T) {
	expected := 7145 // Replace with the actual line count of your test.txt file
	file := utilities.ReadFile(t, "test.txt")
	defer file.Close()

	result, err := utilities.CountLines(file)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d lines, but got %d", expected, result)
	}
}

func TestCountWords(t *testing.T) {
	expected := 58164
	file := utilities.ReadFile(t, "test.txt")
	defer file.Close()

	result, err := utilities.CountWords(file)
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
	file := utilities.ReadFile(t, "test.txt")
	defer file.Close()

	result, err := utilities.CountCharacters(file)
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
	file := utilities.ReadFile(t, "test.txt")
	defer file.Close()

	arr, err := utilities.CountAll(file)
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
