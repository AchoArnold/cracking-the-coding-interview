package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSudokuSolver_Solve(t *testing.T) {
	testData := []struct {
		name    string
		board   [9][9]int
		isValid bool
		result  [9][9]int
	}{
		{
			"valid board",
			[9][9]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			true,
			[9][9]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
		},
		{
			"invalid board",
			[9][9]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{5, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			false,
			[9][9]int{},
		},
	}

	solver := SudokuSolver{}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.isValid {
				result, err := solver.Solve(data.board)
				assert.Nil(t, err)

				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						assert.Equal(t, result[i][j], data.result[i][j])
					}
				}
			} else {
				_, err := solver.Solve(data.board)
				assert.Error(t, err)
			}
		})
	}
}
