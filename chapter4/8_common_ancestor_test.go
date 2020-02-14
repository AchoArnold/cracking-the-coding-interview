package chapter4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParentAwareBinaryTreeNode_CommonAncestor(t *testing.T) {
	t.Run("Test it returns the correct node", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := CreatParentAwareBinarySearchTree(IntToInterfaceSlice(expected))

		assert.Equal(t, 3, tree.CommonAncestor(1, 5).Value.(int))
	})

	t.Run("Test it returns nil when there's no ancestor", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := CreatParentAwareBinarySearchTree(IntToInterfaceSlice(expected))

		assert.Nil(t, tree.CommonAncestor(10, 5))
	})
}
