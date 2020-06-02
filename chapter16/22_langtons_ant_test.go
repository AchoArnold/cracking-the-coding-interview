package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLangtonsAnt(t *testing.T) {
	testData := []struct {
		name   string
		moves  int
		result string
	}{
		{
			"no move",
			0,
			"→\nAnt: →. \n",
		},
		{
			"ten moves",
			10,
			"XX_\nXXX\n←_X\nAnt: ←. \n",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, LangtonsAnt(data.moves))
		})
	}
}
