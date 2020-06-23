package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissing(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		result int
	}{
		{
			"no number missing",
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
			8,
		},
		{
			"odd number missing",
			[]int{0, 1, 2, 4, 5, 6, 7},
			3,
		},
		{
			"even number missing",
			[]int{0, 1, 2, 3, 5, 6, 7},
			4,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, FindMissing(data.input))
		})
	}
}
