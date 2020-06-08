package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand7(t *testing.T) {
	for i := 0; i < 10000; i++ {
		val := Rand7()
		assert.True(t, val < 7)
		assert.True(t, val >= 0)
	}
}
