package chapter16

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionPoint(t *testing.T) {
	testData := []struct {
		name     string
		start1   point
		end1     point
		start2   point
		end2     point
		expected point
	}{
		{
			"2 lines intersect",
			newPoint(6, 4),
			newPoint(2, 3),
			newPoint(1, 6),
			newPoint(3, 2),
			newPoint(2.444, 3.111),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := getIntersectionPoint(data.start1, data.end1, data.start2, data.end2)
			assert.True(t, decimal.NewFromFloat(data.expected.x).Round(3).Equal(decimal.NewFromFloat(actual.x).Round(3)))
			assert.True(t, decimal.NewFromFloat(data.expected.y).Round(3).Equal(decimal.NewFromFloat(actual.y).Round(3)))
		})
	}

	t.Run("2 lines on the same slope with common middle point", func(t *testing.T) {
		actual := getIntersectionPoint(newPoint(1, 1), newPoint(4, 4), newPoint(3, 3), newPoint(6, 6))
		assert.True(t, decimal.NewFromFloat(3).Round(3).Equal(decimal.NewFromFloat(actual.x).Round(3)))
		assert.True(t, decimal.NewFromFloat(3).Round(3).Equal(decimal.NewFromFloat(actual.y).Round(3)))
	})

	t.Run("parallel lines that never meet on the same slope", func(t *testing.T) {
		assert.Nil(t, getIntersectionPoint(newPoint(1, 1), newPoint(4, 4), newPoint(10, 10), newPoint(6, 6)))
	})
}
