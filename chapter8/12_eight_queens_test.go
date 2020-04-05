package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceNQueens(t *testing.T) {
	t.Run("Testing 8 queens", func(t *testing.T) {
		assert.Equal(t, 92, PlaceNQueens(8))
	})

	t.Run("Testing 8 queens", func(t *testing.T) {
		assert.Equal(t, 73712, PlaceNQueens(13))
	})
}
