package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMajorityElement(t *testing.T) {
	testData := []struct {
		name   string
		input  []int
		result int
	}{
		{
			"empty array",
			[]int{},
			-1,
		},
		{
			"one element",
			[]int{5},
			5,
		},
		{
			"mixed list",
			[]int{1, 2, 5, 9, 5, 9, 5, 5, 5},
			5,
		},
		{
			"homogenious list",
			[]int{4, 4, 4, 4},
			4,
		},
		{
			"k is last but not majority",
			[]int{2, 3, 4, 5, 5},
			-1,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, FindMajorityElement(data.input))
		})
	}
}
