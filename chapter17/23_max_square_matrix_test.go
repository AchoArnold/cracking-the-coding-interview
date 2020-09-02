package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSquare(t *testing.T) {
	testData := []struct {
		name   string
		matrix [][]int
		result *SubSquare
	}{
		{
			"no subsquare",
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
			},
			nil,
		},
		{
			"no subsquare",
			[][]int{
				{1, 0, 0},
				{1, 0, 0},
			},
			&SubSquare{0, 1, 1},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.result == nil {
				assert.Nil(t, FindSquare(data.matrix))
			} else {
				val := FindSquare(data.matrix)
				assert.Equal(t, data.result.row, val.row)
				assert.Equal(t, data.result.col, val.col)
				assert.Equal(t, data.result.squareSize, val.squareSize)
			}
		})
	}
}
