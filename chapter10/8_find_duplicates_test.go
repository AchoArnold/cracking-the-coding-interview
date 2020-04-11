package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDuplicates(t *testing.T) {
	t.Run("Testing when there are no duplicates", func(t *testing.T) {
		test := []int{
			1,
			2,
			4,
			5,
			6,
		}
		assert.Equal(t, []int{}, GetDuplicates(test))
	})

	t.Run("Testing when value is not available", func(t *testing.T) {
		test := []int{
			1,
			2,
			4,
			5,
			6,
			6,
			6,
			2,
		}
		assert.Equal(t, []int{6, 6, 2}, GetDuplicates(test))
	})

}
