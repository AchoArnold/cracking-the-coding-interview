package chapter5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawLine(t *testing.T) {
	var testScreen = make([]byte, 8)

	DrawLine(testScreen, 8, 0, 4, 3)

	assert.Equal(t, byte(0x10), testScreen[3])
}
