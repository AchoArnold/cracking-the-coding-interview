package chapter4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
	expected := []int{1, 2,3,4,5}

	tree := &BinaryTreeNode{
		Value: 3,
		Left:  &BinaryTreeNode{
			Value: 2,
			Right:  nil,
			Left: &BinaryTreeNode{
				Value: 1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTreeNode{
			Value: 4,
			Left:  nil,
			Right: &BinaryTreeNode{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, IntToInterfaceSlice(expected), DepthFirstSearchBinaryTree(tree))
}
