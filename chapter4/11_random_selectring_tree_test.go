package chapter4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomTreeNode(t *testing.T) {
	node := NewRandomTreeNode(4)

	assert.Equal(t, 4, node.Value)
	assert.Nil(t, node.Left)
	assert.Nil(t, node.Right)
	assert.Equal(t, 1, node.Size)
}

func TestRandomTreeNode_GetRandomNode(t *testing.T) {
	node := NewRandomTreeNode(5)
	for i, j := 5, 4; i <= 10; i++ {
		node.InsertInOrder(i)
		node.InsertInOrder(j)
		j--
	}

	for i := 0; i <= 10; i++ {
		t.Run("Getting a random node", func(t *testing.T) {
			node := node.GetRandomNode()
			assert.True(t, node.Value >= 0 && node.Value <= 10)
		})
	}

}
