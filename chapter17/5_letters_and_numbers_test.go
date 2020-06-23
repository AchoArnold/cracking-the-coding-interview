package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestSubarray(t *testing.T) {
	testData := []struct {
		name   string
		input  []rune
		result []rune
	}{
		{
			"2 values",
			[]rune{'1', 'a'},
			[]rune{'1', 'a'},
		},
		{
			"more values",
			[]rune("aaaa11a11aa1aa1aaaaa"),
			[]rune("a11a11aa1aa1"),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.result, FindLongestSubarray(data.input))
		})
	}
}
