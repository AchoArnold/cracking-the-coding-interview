package chapter5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitSwapRequired(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		assert.Equal(t, 2, BitSwapRequired(29, 15))
	})

	t.Run("same", func(t *testing.T) {
		assert.Equal(t, 0, BitSwapRequired(29, 29))
	})
}

func TestBitSwapRequiredLastSignificantBit(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		assert.Equal(t, 2, BitSwapRequiredLastSignificantBit(29, 15))
	})

	t.Run("same", func(t *testing.T) {
		assert.Equal(t, 0, BitSwapRequiredLastSignificantBit(29, 29))
	})
}
