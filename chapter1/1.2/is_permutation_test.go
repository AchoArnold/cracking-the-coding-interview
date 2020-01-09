package main

import "testing"

func TestIsPermutationReturnsTrueForStringsWhichArePermutations(test *testing.T) {
	testStrings := map[string]string{
		"abc": "abc",
		"": "",
		"a": "a",
	}

	test.Run("Using brute force", func (test *testing.T) {
		for inputString1, inputString2 := range testStrings {
			if !isPermutation(inputString1, inputString2) {
				test.Errorf("Flase Negative = %s and %s", inputString1, inputString2)
			}
		}
	})
}

func TestIsPermutationReturnsFalseForStringsWhichAreNotPermutations(test *testing.T) {
	testStrings := map[string]string{
		"abcd": "abc",
		"dddadfdsf": "yddadfdsf",
		"": "a",
	}

	test.Run("Using brute force", func (test *testing.T) {
		for inputString1, inputString2 := range testStrings {
			if isPermutation(inputString1, inputString2) {
				test.Errorf("False Possitive = %s and %s", inputString1, inputString2)
			}
		}
	})
}
