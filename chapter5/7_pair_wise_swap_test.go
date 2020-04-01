package chapter5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapOddEvenBits(t *testing.T) {
	assert.Equal(t, 4, SwapOddEvenBits(8))
}
