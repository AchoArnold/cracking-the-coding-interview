package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountEval(t *testing.T) {
	t.Run("Counts the number of parenthesis in the default expression", func(t *testing.T) {
		assert.Equal(t, 5, CountEval("1^0|0|1", false))
	})
}
