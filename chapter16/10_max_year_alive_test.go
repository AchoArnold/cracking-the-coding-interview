package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAliveYear(t *testing.T) {
	testData := []struct {
		name           string
		people         []person
		min, max, year int
	}{
		{
			"all born on the same year",
			[]person{
				{birth: 1900, death: 1950},
				{birth: 1900, death: 1950},
				{birth: 1900, death: 1950},
			},
			1900,
			2000,
			1900,
		},

		{
			"different years",
			[]person{
				{birth: 1900, death: 1950},
				{birth: 1920, death: 1940},
				{birth: 1932, death: 1960},
			},
			1900,
			2000,
			1932,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.year, MaxAliveYear(data.people, data.min, data.max))
		})
	}
}
