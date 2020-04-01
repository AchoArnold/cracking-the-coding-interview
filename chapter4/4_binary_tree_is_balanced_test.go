package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeNode_IsBalanced(t *testing.T) {
	t.Run("It returns true when the tree is balanced", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := GetBinaryTreeFromSortedArray(expected)

		assert.True(t, tree.IsBalanced())
	})

	t.Run("It returns false when the tree is not balanced", func(t *testing.T) {
		tree := BinaryTreeNode{
			Value: 3,
			Left: &BinaryTreeNode{
				Value: 2,
				Left: &BinaryTreeNode{
					Value: 1,
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
		assert.False(t, tree.IsBalanced())
	})
}

func TestBinaryTreeNode_IsBalancedUsingHeightCheck(t *testing.T) {
	t.Run("It returns true when the tree is balanced", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := GetBinaryTreeFromSortedArray(expected)

		assert.True(t, tree.IsBalancedUsingHeightCheck())
	})

	t.Run("It returns false when the tree is not balanced", func(t *testing.T) {
		tree := BinaryTreeNode{
			Value: 3,
			Left: &BinaryTreeNode{
				Value: 2,
				Left: &BinaryTreeNode{
					Value: 1,
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
		assert.False(t, tree.IsBalancedUsingHeightCheck())
	})
}
