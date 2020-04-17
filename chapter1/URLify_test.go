package chapter1

import "testing"

func TestURLify(t *testing.T) {
	inputTestData := []struct {
		inputString    string
		length         int
		expectedOutput string
	}{
		{
			"how are you",
			11,
			"how%20are%20you",
		},
		{
			"how are you           ",
			11,
			"how%20are%20you",
		},
	}

	for _, testInput := range inputTestData {
		t.Run("Testing that it produces the correct results", func(t *testing.T) {
			if URLify(testInput.inputString, testInput.length) != testInput.expectedOutput {
				t.Errorf(
					"Unexpected result\nInput = %s\nLength = %d\nOutput = %s",
					testInput.inputString,
					testInput.length,
					testInput.expectedOutput,
				)
			}
		})
	}
}
