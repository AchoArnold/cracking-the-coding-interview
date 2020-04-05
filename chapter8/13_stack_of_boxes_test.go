package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStackHeight(t *testing.T) {
	t.Run("max height for 1 box", func(t *testing.T) {
		boxes := []Box{
			{2, 2, 3},
		}

		assert.Equal(t, 2, MaxStackHeight(boxes))
	})

	t.Run("max height for 4 boxes", func(t *testing.T) {
		boxes := []Box{
			{5, 2, 3},
			{4, 2, 3},
			{3, 2, 3},
			{2, 2, 3},
		}

		assert.Equal(t, 5, MaxStackHeight(boxes))
	})

	t.Run("max height for 4 boxes", func(t *testing.T) {
		boxes := []Box{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
			{2, 2, 2},
		}

		assert.Equal(t, 6, MaxStackHeight(boxes))
	})
}
