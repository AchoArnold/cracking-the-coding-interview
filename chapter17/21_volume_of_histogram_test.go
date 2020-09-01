package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHistogramVolume(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		output int
	}{
		{
			"no item",
			[]int{},
			0,
		},
		{
			"a couple of items",
			[]int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 5, 0, 1, 0, 0, 0},
			26,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.output, ComputeHistogramVolume(data.input))
		})
	}
}
