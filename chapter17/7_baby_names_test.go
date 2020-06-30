package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrulyMostPopular(t *testing.T) {
	testData := []struct {
		name          string
		inputNames    map[string]int
		inputSynonyms [][]string
		result        []int
	}{
		{
			"valid test data",
			map[string]int{
				"John":     10,
				"Jon":      3,
				"Davis":    2,
				"Kari":     3,
				"Johnny":   11,
				"Carlton":  8,
				"Carleton": 2,
				"Jonathan": 9,
				"Carrie":   5,
			},
			[][]string{
				{"Jonathan", "John"},
				{"Jon", "Johnny"},
				{"Johnny", "John"},
				{"Kari", "Carrie"},
				{"Carleton", "Carlton"},
			},
			[]int{2, 10, 33, 8},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := TrulyMostPopular(data.inputNames, data.inputSynonyms)
			var actualInt []int
			for _, freq := range actual {
				actualInt = append(actualInt, freq)
			}

			assert.ElementsMatch(t, data.result, actualInt)
		})
	}
}
