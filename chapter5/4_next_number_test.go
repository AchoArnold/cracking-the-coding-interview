package chapter5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextNumberWithSame1Bits(t *testing.T) {
	assert.Equal(t, 1401, GetNextNumberWithSame1Bits(1398))
}

func TestGetPreviousNumberWithSame1Bits(t *testing.T) {
	assert.Equal(t, 1397, GetPreviousNumberWithSame1Bits(1398))
}
