package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElementInSortedMatrix(t *testing.T) {
	testMatrix := [][]int{
		{15, 20, 70, 85},
		{20, 35, 80, 95},
		{30, 55, 94, 105},
	}

	t.Run("Testing when there are no duplicates", func(t *testing.T) {
		result := FindElementInSortedMatrix(testMatrix, 95)

		expected := MatrixCoordinate{1, 3}

		assert.Equal(t, expected.row, result.row)
		assert.Equal(t, expected.column, result.column)
	})

	t.Run("Testing when value is not available", func(t *testing.T) {
		assert.Nil(t, FindElementInSortedMatrix(testMatrix, 200))
	})
}
