package chapter4

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeNode_ContainsTree(t *testing.T) {
	t.Run("Test it returns true when the child is nil", func(t *testing.T) {
		parentTree := GetBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
		childTree := GetBinaryTreeFromSortedArray([]int{})

		assert.True(t, parentTree.ContainsTree(childTree))
	})

	t.Run("Test it returns true when the root contains the child", func(t *testing.T) {
		parentTree := GetBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
		childTree := GetBinaryTreeFromSortedArray([]int{1, 2})

		assert.True(t, parentTree.ContainsTree(childTree))
	})

	t.Run("Test it returns false when the does not contain the child", func(t *testing.T) {
		parentTree := GetBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
		childTree := GetBinaryTreeFromSortedArray([]int{4, 3, 5})

		assert.False(t, parentTree.ContainsTree(childTree))
	})
}
