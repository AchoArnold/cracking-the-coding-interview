package main

import (
	"testing"

	util "github.com/AchoArnold/cracking-the-coding-interview/utils"
)

func TestOneAway(t *testing.T) {
	testData := []struct {
		string1, string2 string
		expectedResult   bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"paless", "pale", false},
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", true},
		{" ", " ", true},
		{"", "", true},
	}

	for _, data := range testData {
		t.Run("running tests", func(t *testing.T) {
			if oneAway(data.string1, data.string2) != data.expectedResult {
				t.Errorf(
					"False %s: \"%s\" and \"%s\" is %sone away",
					util.IfThenElse(data.expectedResult, "negative", "positive"),
					data.string1,
					data.string2,
					util.IfThenElse(data.expectedResult, "", "not "),
				)
			}
		})
	}
}
