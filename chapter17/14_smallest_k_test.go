package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestK(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		k      int
		output []int
	}{
		{
			"k larger than elements",
			[]int{1, 2, 4},
			4,
			[]int{},
		},
		{
			"k less than zero",
			[]int{1, 2, 4},
			-1,
			[]int{},
		},
		{
			"valid k",
			[]int{9, 23, 5, 1, 4, 92, 5, 2},
			4,
			[]int{1, 2, 4, 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.output, SmallestK(data.input, data.k))
		})
	}
}
