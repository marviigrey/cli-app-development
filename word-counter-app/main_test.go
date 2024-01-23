package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	// Test case for counting words
	t.Run("Count Words", func(t *testing.T) {
		input := "This is a test sentence."
		reader := bytes.NewReader([]byte(input))
		countLines := false
		countBytes := false

		result := count(reader, countLines, countBytes)
		expected := 5 // Assuming the sentence has 5 words

		if result != expected {
			t.Errorf("Expected %d words, but got %d", expected, result)
		}
	})

	// Test case for counting lines
	t.Run("Count Lines", func(t *testing.T) {
		input := "Line 1\nLine 2\nLine 3"
		reader := bytes.NewReader([]byte(input))
		countLines := true
		countBytes := false

		result := count(reader, countLines, countBytes)
		expected := 3 // Assuming there are 3 lines

		if result != expected {
			t.Errorf("Expected %d lines, but got %d", expected, result)
		}
	})

	// Test case for counting bytes
	t.Run("Count Bytes", func(t *testing.T) {
		input := "Hello, World!"
		reader := bytes.NewReader([]byte(input))
		countLines := false
		countBytes := true

		result := count(reader, countLines, countBytes)
		expected := len(input) // Expected result is the length of the input string

		if result != expected {
			t.Errorf("Expected %d bytes, but got %d", expected, result)
		}
	})
}

// You can add more test cases for different scenarios.
