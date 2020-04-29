package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUnsortedSequence(t *testing.T) {
	testData := []struct {
		name  string
		array []int
		left  int
		right int
	}{
		{
			"sorted array",
			[]int{1, 2, 3, 4, 5, 6, 7},
			-1,
			-1,
		},
		{
			"unsorted array",
			[]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19},
			3,
			9,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			left, right := FindUnsortedSequence(data.array)
			assert.Equal(t, data.right, right)
			assert.Equal(t, data.left, left)
		})
	}
}
