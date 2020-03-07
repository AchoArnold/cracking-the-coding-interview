package chapter5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapOddEvenBits(t *testing.T) {
	assert.Equal(t, 4, SwapOddEvenBits(8))
}
