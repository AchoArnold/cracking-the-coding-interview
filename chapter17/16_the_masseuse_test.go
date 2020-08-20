package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	name           string
	input          []int
	expectedOutput int
}{
	{
		"Valid sequence",
		[]int{30, 15, 60, 75, 45, 15, 15, 45},
		180,
	},
	{
		"no item",
		[]int{},
		0,
	},
}

func TestMaxMinutesRecursive(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedOutput, MaxMinutesRecursive(data.input))
		})
	}
}

func TestMaxMinutesIterative(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedOutput, MaxMinutesIterative(data.input))
		})
	}
}
