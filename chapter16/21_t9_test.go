package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidT9Words(t *testing.T) {
	testData := []struct {
		name       string
		input      string
		dictionary map[string]string
		results    []string
	}{
		{
			"valid test case",
			"8733",
			map[string]string{
				"how":  "how",
				"fine": "fine",
				"tree": "tree",
				"used": "used",
			},
			[]string{"tree", "used"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.results, GetValidT9Words(data.input, data.dictionary))
		})
	}
}
