package chapter8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var invalidSet = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var validIndex = 60000

func TestMagicFirstBruteForce(t *testing.T) {
	t.Run("Test it returns -1 if value is not found", func(t *testing.T) {
		assert.Equal(t, -1, MagicFirstBruteForce(invalidSet))
	})

	t.Run("test it returns the valid index if x is valid", func(t *testing.T) {
		assert.Equal(t, validIndex, MagicFirstBruteForce(getValidSet()))
	})
}

func TestMagicBinarySearch(t *testing.T) {
	t.Run("Test it returns -1 if value is not found", func(t *testing.T) {
		assert.Equal(t, -1, MagicBinarySearch(invalidSet))
	})

	t.Run("test it returns the valid index if x is valid", func(t *testing.T) {
		assert.Equal(t, validIndex, MagicBinarySearch(getValidSet()))
	})
}

func BenchmarkMagicFirstBruteForce(b *testing.B) {
	val := getValidSet()
	for i := 0; i < b.N; i++ {
		MagicFirstBruteForce(val)
	}
}

func BenchmarkMagicBinarySearch(b *testing.B) {
	val := getValidSet()
	for i := 0; i < b.N; i++ {
		MagicBinarySearch(val)
	}
}

func getValidSet() []int {
	var validSet [100001]int
	for i := -1; i < 100000; i++ {
		validSet[i+1] = i
	}

	validSet[validIndex] = validIndex

	return validSet[:]
}
