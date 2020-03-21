package chapter8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPowerSet(t *testing.T) {
	set := []int{1, 2, 3, 4}

	t.Run("Get all combinations", func(t *testing.T) {
		powerSet := GetPowerSet(set)
		assert.Equal(t, 16, len(powerSet))
	})
}