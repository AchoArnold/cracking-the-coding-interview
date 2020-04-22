package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFactorialZeros(t *testing.T) {
	testData := []struct {
		name           string
		n              int
		expectedResult int
	}{
		{"has 1 zero", 5, 1},
		{"has no zero", 3, 0},
		{"has 6 zeros", 25, 6},
		{"returns -1 if num is invalid", -32, -1},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedResult, CountFactorialZeros(data.n))
		})
	}
}
