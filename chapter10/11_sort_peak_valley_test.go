package chapter10

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPeakValley(t *testing.T) {
	array := []int{9, 1, 0, 4, 8}

	SortPeakValley(array)

	assert.True(t, reflect.DeepEqual([]int{1, 9, 0, 8, 4}, array))
}

func TestSwapArray(t *testing.T) {
	array := []int{9, 1, 0, 4, 8}

	SwapArray(array, 0, 4)

	assert.Equal(t, []int{8, 1, 0, 4, 9}, array)

	SwapArray(array, 0, 4)

	assert.Equal(t, []int{9, 1, 0, 4, 8}, array)
}
