package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentAwareBinaryTreeNode_InOrderSuccessor(t *testing.T) {
	t.Run("It returns the in order successor when there is a successor", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}

		tree := CreatParentAwareBinarySearchTree(IntToInterfaceSlice(expected))

		assert.Equal(t, 2, tree.Left.Left.InOrderSuccessor().Value.(int))
	})

	t.Run("It returns nil when there is no in order successor when there is a successor", func(t *testing.T) {
		expected := []int{1}

		tree := CreatParentAwareBinarySearchTree(IntToInterfaceSlice(expected))

		assert.Nil(t, tree.InOrderSuccessor())
	})
}
