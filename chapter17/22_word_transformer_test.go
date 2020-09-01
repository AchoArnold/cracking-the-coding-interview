package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformWord(t *testing.T) {
	testData := []struct {
		name       string
		dictionary []string
		source     string
		target     string
		output     []string
	}{
		//{
		//	"no value",
		//	[]string{},
		//	"one",
		//	"two",
		//	[]string{},
		//},
		{
			"dictionary with values",
			[]string{"but", "put", "big", "pot", "pog", "dog", "lot"},
			"bit",
			"dog",
			[]string{"bit", "but", "put", "pot", "pog", "dog"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.output, TransformWord(data.dictionary, data.source, data.target))
		})
	}
}
