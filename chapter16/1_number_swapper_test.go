package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapNumbersInPlace(t *testing.T) {
	t.Run("swap works when numbers are the same", func(t *testing.T) {
		a := 1
		b := 1

		SwapNumbersInPlace(&a, &b)

		assert.Equal(t, 1, a)
		assert.Equal(t, 1, b)
	})

	t.Run("swap works when the numbers are different", func(t *testing.T) {
		a := 432
		b := -122

		SwapNumbersInPlace(&a, &b)

		assert.Equal(t, -122, a)
		assert.Equal(t, 432, b)
	})
}
