package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripleStep(t *testing.T) {
	t.Run("It returns zeo if n <= 0", func(t *testing.T) {
		assert.Equal(t, 0, TripleStep(0))
		assert.Equal(t, 0, TripleStep(-1))
	})

	t.Run("It returns 1 if n is between 1 and 3", func(t *testing.T) {
		assert.Equal(t, 1, TripleStep(1))
		assert.Equal(t, 2, TripleStep(2))
		assert.Equal(t, 4, TripleStep(3))
	})

	t.Run("It returns the correct number of steps", func(t *testing.T) {
		assert.Equal(t, 927, TripleStep(12))
	})
}

func TestTripleStepRecursive(t *testing.T) {
	t.Run("It returns zeo if n <= 0", func(t *testing.T) {
		assert.Equal(t, 0, TripleStepRecursive(0))
		assert.Equal(t, 0, TripleStepRecursive(-1))
	})

	t.Run("It returns 1 if n is between 1 and 3", func(t *testing.T) {
		assert.Equal(t, 1, TripleStepRecursive(1))
		assert.Equal(t, 2, TripleStepRecursive(2))
		assert.Equal(t, 4, TripleStepRecursive(3))
	})

	t.Run("It returns the correct number of steps", func(t *testing.T) {
		assert.Equal(t, 927, TripleStepRecursive(12))
	})
}

func BenchmarkTripleStep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TripleStep(15)
	}
}

func BenchmarkTripleStepRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TripleStepRecursive(15)
	}
}
