package chapter8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPermutations(t *testing.T) {
	t.Run("it returns an empty array when there's no string", func(t *testing.T) {
		assert.Equal(t, []string{""}, GetPermutations(""))
	})

	t.Run("It returns an array of all possible combinations", func(t *testing.T) {
		assert.Equal(t, 24, len(GetPermutations("fine")))
	})
}
