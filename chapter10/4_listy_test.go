package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchListy(t *testing.T) {
	list := Listy{
		0: 2,
		1: 3,
		2: 8,
		3: 20,
		4: 21,
		5: 21,
		6: 23,
	}

	t.Run("Testing when value is available", func(t *testing.T) {
		assert.Equal(t, 6, SearchListy(list, 23))
	})

	t.Run("Testing when value is not available", func(t *testing.T) {
		assert.Equal(t, -1, SearchListy(list, 40))
	})
}
