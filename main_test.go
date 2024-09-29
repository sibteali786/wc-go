package main

import (
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
	result, err := utilities.CountLines("test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
