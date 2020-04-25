package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumWithoutComparison(t *testing.T) {
	testData := []struct {
		name      string
		a, b, max int
	}{
		{"comparing possitive numbers", 1, 2, 2},
		{"comparing zero values", 0, 0, 0},
		{"comparing negative numbers", -123, -1212, -123},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.max, MaxNumWithoutComparison(data.a, data.b))
		})
	}
}
