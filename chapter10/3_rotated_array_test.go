package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotatedArray(t *testing.T) {
	t.Run("Testing when value is available", func(t *testing.T) {
		testPart1 := []int{50, 2, 5, 5, 20, 30, 40}

		assert.Equal(t, 1, SearchRotatedArray(testPart1, 2))
	})

	t.Run("Testing when value is not available", func(t *testing.T) {
		testPart1 := []int{50, -1, 2, 5, 5, 20, 30, 40}

		assert.Equal(t, -1, SearchRotatedArray(testPart1, 4))
	})
}
