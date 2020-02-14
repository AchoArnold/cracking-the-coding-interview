package chapter4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeNode_IsBinarySearchTree(t *testing.T) {
	t.Run("It returns true when the tree is a valid binary search tree", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := GetBinaryTreeFromSortedArray(expected)

		assert.True(t, tree.IsBinarySearchTree())
	})

	t.Run("It returns false when the tree is not a valid binary search tree", func(t *testing.T) {
		tree := BinaryTreeNode{
			Value: 3,
			Left: &BinaryTreeNode{
				Value: 2,
				Left: &BinaryTreeNode{
					Value: 6,
					Left: &BinaryTreeNode{
						Value: 0,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: &BinaryTreeNode{
				Value: 4,
				Left:  nil,
				Right: nil,
			},
		}
		assert.False(t, tree.IsBinarySearchTree())
	})
}
