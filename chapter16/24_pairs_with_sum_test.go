package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairsWithSum(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		sum    int
		result [][]int
	}{
		{
			"empty input",
			[]int{},
			0,
			[][]int{},
		},
		{
			"negative and positive",
			[]int{1, 2, 3, -2, 9, 10, -1, 0},
			0,
			[][]int{
				{-1, 1},
				{-2, 2},
			},
		},
		{
			"sum does not exist",
			[]int{1, 2, 3, -2, 9, 10, -1, 0},
			20,
			[][]int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.result, PairsWithSum(data.input, data.sum))
		})
	}
}
