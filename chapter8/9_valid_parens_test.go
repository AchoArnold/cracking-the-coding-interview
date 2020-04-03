package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAllValidParens(t *testing.T) {
	t.Run("Test it generates paranthesis when count = 0", func(t *testing.T) {
		assert.Equal(t, []string{""}, GenerateAllValidParens(0))
	})

	t.Run("Test it generates paranthesis when count = 1", func(t *testing.T) {
		assert.Equal(t, []string{"()"}, GenerateAllValidParens(1))
	})

	t.Run("Test it generates paranthesis when count = 3", func(t *testing.T) {
		assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, GenerateAllValidParens(3))
	})
}
