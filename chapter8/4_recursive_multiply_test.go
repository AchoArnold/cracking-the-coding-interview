package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveMultiply(t *testing.T) {
	t.Run("Test multiplication by 0", func(t *testing.T) {
		assert.Equal(t, 0, RecursiveMultiply(3, 0))
		assert.Equal(t, 0, RecursiveMultiply(0, 33))
	})

	t.Run("Test m x n multiplication", func(t *testing.T) {
		assert.Equal(t, 80, RecursiveMultiply(8, 10))
	})
}
