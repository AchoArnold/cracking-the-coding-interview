package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutationsWithDuplicates(t *testing.T) {
	t.Run("it returns an empty array when there's no string", func(t *testing.T) {
		assert.Equal(t, []string{""}, GetPermutationsWithDuplicates(""))
	})

	t.Run("It returns an array of all possible combinations", func(t *testing.T) {
		assert.Equal(t, 24, len(GetPermutationsWithDuplicates("fine")))
	})

	t.Run("It returns a single array when the string contains only repeating characters", func(t *testing.T) {
		assert.Equal(t, 1, len(GetPermutationsWithDuplicates("aaaaaaaaa")))
	})

	t.Run("It returns a valid permutation when there are repeating characters", func(t *testing.T) {
		assert.Equal(t, 3, len(GetPermutationsWithDuplicates("aab")))
	})
}
