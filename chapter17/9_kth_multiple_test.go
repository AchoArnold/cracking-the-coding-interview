package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKthMagicNumber(t *testing.T) {
	testData := []struct {
		name   string
		k      int
		result int
	}{
		{
			"k is 5",
			5,
			81,
		},
		{
			"k is 1",
			1,
			1,
		},
		{
			"k is 13",
			13,
			531441,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, GetKthMagicNumber(data.k))
		})
	}
}
