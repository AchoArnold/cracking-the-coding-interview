package chapter8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaintFill(t *testing.T) {
	t.Run("it returns false on an empty screen", func(t *testing.T) {
		assert.False(t, PaintFill([][]Color{{}}, 0, 0, Color("Green")))
	})

	t.Run("it returns a correctly painted screen screen", func(t *testing.T) {
		colors := map[int]Color{
			0: Color("Blue"),
			1: Color("Green"),
		}

		screen := make([][]Color, 5)
		for i := 0; i < 5; i++ {
			screen[i] = make([]Color, 5)
			for j := 0; j < 5; j++ {
				screen[i][j] = colors[j%2]
			}
		}

		assert.True(t, PaintFill(screen, 2, 2, colors[1]))
	})
}
