package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBinaryTreeFromSortedArray(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := GetBinaryTreeFromSortedArray(expected)

	assert.Equal(t, IntToInterfaceSlice(expected), DepthFirstSearchBinaryTree(tree))
}
