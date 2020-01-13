package main

import (
	util "github.com/AchoArnold/cracking-the-coding-interview/utils"
	"testing"
)

func TestOneAway(t *testing.T) {
	testData :=[]struct {
		string1 string
		string2 string
		expectedResult bool
	}{
		{
			string1: "pale",
			string2: "ple",
			expectedResult: true,
		},
		{
			string1: "pales",
			string2: "pale",
			expectedResult: true,
		},
		{
			string1: "pale",
			string2: "bale",
			expectedResult: true,
		},
		{
			string1: "pale",
			string2: "bake",
			expectedResult: false,
		},
		{
			string1: "paless",
			string2: "pale",
			expectedResult: false,
		},
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