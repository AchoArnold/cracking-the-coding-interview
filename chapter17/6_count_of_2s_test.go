package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount2sInRange(t *testing.T) {
	testData := []struct {
		name   string
		input  int
		result int
	}{
		{
			"zero",
			0,
			0,
		},
		{
			"one",
			0,
			0,
		},
		{
			"12",
			12,
			2,
		},
		{
			"61523",
			61523,
			34507,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, Count2sInRange(data.input))
		})
	}
}
