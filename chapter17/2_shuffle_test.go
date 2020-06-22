package chapter17

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffleArray(t *testing.T) {
	testData := []struct {
		name  string
		input []int
	}{
		{
			"sequential array",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			input := append([]int{}, data.input...)
			ShuffleArray(data.input)

			assert.False(t, reflect.DeepEqual(input, data.input))
		})
	}
}
