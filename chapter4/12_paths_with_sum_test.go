package chapter4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeNode_CountPathsWithSum(t *testing.T) {
	t.Run("Test it gives the correct path when paths exist", func(t *testing.T) {
		root := GetBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		assert.Equal(t, 2, root.CountPathsWithSum(13))
	})

	t.Run("Test it gives the correct path when paths do not exist", func(t *testing.T) {
		root := GetBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		assert.Equal(t, 0, root.CountPathsWithSum(60))
	})
}
