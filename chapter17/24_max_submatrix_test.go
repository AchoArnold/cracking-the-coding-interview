package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMatrix(t *testing.T) {
	testData := []struct {
		name   string
		matrix [][]int
		result *SubMatrix
	}{
		{
			"empty matrix",
			[][]int{{}},
			nil,
		},
		{
			"valid matrix",
			[][]int{
				{0, 1, -2, 1},
				{-10, -5, 2, 6},
				{4, 5, 1, 0},
			},
			&SubMatrix{
				row1: 2,
				col1: 0,
				row2: 2,
				col2: 2,
				sum:  10,
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.result == nil {
				assert.Nil(t, MaxMatrix(data.matrix))
			} else {
				val := MaxMatrix(data.matrix)
				assert.Equal(t, data.result.row1, val.row1)
				assert.Equal(t, data.result.col1, val.col1)
				assert.Equal(t, data.result.row2, val.row2)
				assert.Equal(t, data.result.col2, val.col2)
				assert.Equal(t, data.result.sum, val.sum)
			}
		})
	}
}
