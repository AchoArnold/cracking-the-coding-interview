package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningMedianCalculator_Median(t *testing.T) {
	testData := []struct {
		name   string
		stream []int
		median []float64
	}{
		{
			"one item",
			[]int{5},
			[]float64{5},
		},
		{
			"5 items in order",
			[]int{1, 2, 3, 4, 5},
			[]float64{1, 1.5, 2, 2.5, 3},
		},
		{
			"5 items out of order",
			[]int{1, 3, 5, 2, 4},
			[]float64{1, 2, 3, 2.5, 3},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			calculator := NewRunningMedianCalculator()

			for i := 0; i < len(data.stream); i++ {
				calculator.addItem(data.stream[i])
				assert.Equal(t, data.median[i], calculator.Median())
			}
		})
	}
}
