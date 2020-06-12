package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Compute(t *testing.T) {
	testData := []struct {
		name  string
		input string
		value float64
	}{
		{
			"no value",
			"",
			float64(0),
		},
		{
			"no value",
			"2*3+5/6*3+15",
			23.5,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			calc := newCalculator()
			assert.Equal(t, data.value, calc.Compute(data.input))
		})
	}
}
