package main

import (
	"testing"
)

func TestStringCompression(t *testing.T) {
	inputOuptut := map[string]string{
		"aabcccccaaa" : "a2b1c5a3",
		"abcdef": "abcdef",
	}

	for input, output := range inputOuptut {
		t.Run("Testing String", func(t *testing.T) {
			if output != stringCompression(input) {
				t.Errorf("Input %s did not produce the output %s", input, output)
			}
		})

	}
}
