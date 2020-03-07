package chapter5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBits(t *testing.T) {
	assert.Equal(t, 1100, SetBitsInRange(1024, 19, 2, 6))
}
