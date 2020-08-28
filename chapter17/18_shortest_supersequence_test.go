package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestSuperSequence(t *testing.T) {
	testData := []struct {
		name   string
		big    []int
		small  []int
		result Range
	}{
		{
			"valid elements",
			[]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7},
			[]int{1, 5, 9},
			Range{7, 10},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result.length(), ShortestSuperSequence(data.big, data.small).length())
		})

	}
}
