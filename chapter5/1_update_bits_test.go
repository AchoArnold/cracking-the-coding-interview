package chapter5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateBits(t *testing.T) {
	assert.Equal(t, 1100, SetBitsInRange(1024, 19, 2, 6))
}
