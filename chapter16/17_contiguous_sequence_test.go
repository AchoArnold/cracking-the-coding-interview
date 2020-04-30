package chapter16

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	name  string
	input []int
	sum   int
}{
	{
		"empty array",
		[]int{},
		math.MinInt32,
	},
	{
		"sequence array",
		[]int{2, -8, 3, -2, 4, -10},
		5,
	},
	{
		"single element",
		[]int{1, 2, -40, 21},
		21,
	},

	{
		"single negative element",
		[]int{-40},
		-40,
	},
}

func TestGetContiguousSequenceBruteForce(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.sum, GetContiguousSequenceBruteForce(data.input))
		})
	}
}

func TestGetContiguousSequence(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.sum, GetContiguousSequence(data.input))
		})
	}
}

func BenchmarkGetContiguousSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetContiguousSequence(testData[1].input)
	}
}

func BenchmarkGetContiguousSequenceBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetContiguousSequenceBruteForce(testData[1].input)
	}
}
