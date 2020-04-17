package chapter1

import (
	"testing"

	util "github.com/AchoArnold/cracking-the-coding-interview/utils"
)

func TestIsStringRotation(t *testing.T) {
	inputOutput := []struct {
		input1         string
		input2         string
		expectedOutput bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"hellomynameis", "nameishellomy", true},
		{"2", "", false},
	}

	for _, data := range inputOutput {
		t.Run("Testing", func(t *testing.T) {
			if isStringRotation(data.input1, data.input2) != data.expectedOutput {
				t.Errorf(
					"False %s: \"%s\" and \"%s\" is %sa rotation",
					util.IfThenElse(data.expectedOutput, "negative", "positive"),
					data.input1,
					data.input2,
					util.IfThenElse(data.expectedOutput, "", "not "),
				)
			}
		})
	}
}
