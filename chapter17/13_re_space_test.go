package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestSplit(t *testing.T) {
	testData := []struct {
		name       string
		input      string
		dictionary map[string]struct{}
		output     string
	}{
		{
			"valid test case",
			"jesslookedjustliketimherbrother",
			map[string]struct{}{
				"looked":  {},
				"just":    {},
				"like":    {},
				"her":     {},
				"brother": {},
			},
			"j e s s looked just like t i m her brother ",
		},
		{
			"empty dictionary",
			"jesslookedjustliketimherbrother",
			map[string]struct{}{},
			"j e s s l o o k e d j u s t l i k e t i m h e r b r o t h e r ",
		},
		{
			"empty string",
			"",
			map[string]struct{}{},
			"",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.output, BestSplit(data.dictionary, data.input))
		})
	}
}
