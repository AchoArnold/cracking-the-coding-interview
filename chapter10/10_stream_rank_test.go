package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankNode(t *testing.T) {
	stream := []int{5, 1, 4, 4, 5, 9, 7, 13, 3}
	root := NewRankNode(stream[0])
	for i := 1; i < len(stream); i++ {
		root.insert(stream[i])
	}

	t.Run("Testing rank method when a rank exists", func(t *testing.T) {
		assert.Equal(t, 3, root.getRank(9))
	})

	t.Run("Testing rank method when a rank does not exist", func(t *testing.T) {
		assert.Equal(t, -1, root.getRank(200))
	})
}
