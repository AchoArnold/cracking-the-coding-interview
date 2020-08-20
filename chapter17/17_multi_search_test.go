package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAllSubstringsInString(t *testing.T) {
	testData := []struct {
		name           string
		big            string
		smalls         []string
		expectedResult map[string][]int
	}{
		{
			"no smalls in string",
			"holland",
			[]string{"go", "github"},
			map[string][]int{
				"go":     {},
				"github": {},
			},
		},
		{
			"valid set of small strings",
			"mississippi",
			[]string{"is", "ppi", "hi", "sis", "i", "ssippi"},
			map[string][]int{
				"is":     {1, 4},
				"ppi":    {8},
				"hi":     {},
				"sis":    {3},
				"i":      {1, 4, 7, 10},
				"ssippi": {5},
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedResult, SearchAllSubstringsInString(data.big, data.smalls))
		})
	}
}
