package chapter1

import "testing"

func TestUniqueString(t *testing.T) {
	testStrings := map[string]bool{
		"abc": true,
		"def": true,
		"113": false,
		"12451": false,
	}

	t.Run("Using brute force", func (t *testing.T) {
		for inputString, expectedResult := range testStrings {
			if actual := isUnique(inputString); actual != expectedResult {
				t.Errorf("Expected unique=%t for %s", expectedResult, inputString)
			}
		}
	})

	t.Run("Using map data structure", func (t *testing.T) {
		for inputString, expectedResult := range testStrings {
			if actual := isUniqueMap(inputString); actual != expectedResult {
				t.Errorf("Expected unique=%t for %s", expectedResult, inputString)
			}
		}
	})
}