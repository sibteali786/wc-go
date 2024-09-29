package main

import (
	"testing"
	"wcgo/utilities"
)

func TestCountBytes(t *testing.T) {
	expected := 342190
	result, err := utilities.CountBytes("test.txt")
	if err != nil {
		t.Errorf(
			"Some %v",
			err,
		)
	}
	if result != expected {
		t.Errorf("Expected %d , got %d", expected, result)
	}
}
