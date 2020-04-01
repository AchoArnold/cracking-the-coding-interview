package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPowerSet(t *testing.T) {
	set := []int{1, 2, 3, 4}

	t.Run("Get all combinations", func(t *testing.T) {
		powerSet := GetPowerSet(set)
		assert.Equal(t, 16, len(powerSet))
	})
}
