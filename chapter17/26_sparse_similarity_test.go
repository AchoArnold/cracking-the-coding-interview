package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSimilarities(t *testing.T) {
	testData := []struct {
		name   string
		input  []Document
		output map[DocPair]float64
	}{
		{
			"no data",
			[]Document{},
			map[DocPair]float64{},
		},
		{
			"valid data",
			[]Document{
				{
					words: []int{14, 15, 100, 9, 3},
					id:    13,
				},
				{
					words: []int{32, 1, 9, 3, 5},
					id:    16,
				},
				{
					words: []int{15, 29, 2, 6, 8, 7},
					id:    19,
				},
				{
					words: []int{7, 10},
					id:    24,
				},
			},
			map[DocPair]float64{
				DocPair{
					doc1: 13,
					doc2: 19,
				}: 0.1,
				DocPair{
					doc1: 13,
					doc2: 16,
				}: 0.25,
				DocPair{
					doc1: 19,
					doc2: 24,
				}: 0.14285714285714285,
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.output, ComputeSimilarities(data.input))
		})
	}
}
