package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeChange(t *testing.T) {
	t.Run("testing with n = 1", func(t *testing.T) {
		assert.Equal(t, 1, MakeChange(1))
	})

	t.Run("testing with n = 100", func(t *testing.T) {
		assert.Equal(t, 242, MakeChange(100))
	})
}
