package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingSequence(t *testing.T) {
	testData := []struct {
		name   string
		input  []HeightWidth
		output []HeightWidth
	}{
		{
			"valid data",
			[]HeightWidth{{65, 100}, {70, 150}, {70, 150}, {56, 90}, {75, 190}, {60, 95}, {68, 110}},
			[]HeightWidth{{56, 90}, {60, 95}, {65, 100}, {68, 110}, {70, 150}, {75, 190}},
		},
		{
			"empty array",
			[]HeightWidth{},
			[]HeightWidth{},
		},
		{
			"one item",
			[]HeightWidth{{65, 100}},
			[]HeightWidth{{65, 100}},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.output, LongestIncreasingSequence(data.input))
		})
	}
}
