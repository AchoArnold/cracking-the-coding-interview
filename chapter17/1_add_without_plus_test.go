package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWithoutPlus(t *testing.T) {
	testData := []struct {
		name   string
		a      int
		b      int
		result int
	}{
		{
			"zero",
			0,
			0,
			0,
		},
		{
			"2 negative",
			-5,
			-3,
			-8,
		},
		{
			"2 possitive",
			5,
			2,
			7,
		},
		{
			"mixed",
			-104,
			6,
			-98,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, AddWithoutPlus(data.a, data.b))
		})
	}

}
