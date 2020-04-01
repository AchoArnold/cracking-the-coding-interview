package chapter5

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSequenceBruteForce(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		assert.Equal(t, 8, LongestSequenceBruteForce(1775))
	})

	t.Run("Max int", func(t *testing.T) {
		assert.Equal(t, 8, LongestSequenceBruteForce(math.MaxInt32))
	})
}

func TestFlipBit(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		assert.Equal(t, 8, FlipBit(1775))
	})

	t.Run("Max int", func(t *testing.T) {
		assert.Equal(t, 32, FlipBit(math.MaxInt32))
	})
}
