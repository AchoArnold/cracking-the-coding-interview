package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSmallestNonNegativeDifference(t *testing.T) {
	array1 := []int{1, 3, 15, 11, 2}
	array2 := []int{23, 127, 235, 19, 8}

	assert.Equal(t, 3, FindSmallestNonNegativeDifference(array1, array2))
}
