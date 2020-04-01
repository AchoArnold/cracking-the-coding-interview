package main

import "testing"

func TestPalindromePermutation(t *testing.T) {
	testData := map[string]bool{
		"cdefghmnopqrstuvw":    false,
		"cdef ghmn opq rstuvw": false,
		"cdcdcdceeseef":        true,
		"cd cdcdcdeeeef":       true,
	}

	for inputString, expectedResult := range testData {
		t.Run("TestCase", func(t *testing.T) {
			if isPalindromePermutation(inputString) != expectedResult {
				t.Errorf(
					"False %s: %s is %sa plindrome",
					IfThenElse(expectedResult, "negative", "positive"),
					inputString,
					IfThenElse(expectedResult, "", "not "),
				)
			}
		})
	}

}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
