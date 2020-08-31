package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingTwo(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			"no item",
			[]int{},
			[]int{1, 2},
		},
		{
			"items with none missing",
			[]int{1, 2, 3},
			[]int{4, 5},
		},
		{
			"items with two missing",
			[]int{1, 2, 3, 5, 6, 8, 9},
			[]int{4, 7},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.output, MissingTwo(data.input))
		})
	}
}
