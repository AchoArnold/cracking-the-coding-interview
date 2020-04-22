package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasWonTicTacToe(t *testing.T) {
	blue := 1
	red := 2
	empty := 0

	testData := []struct {
		name   string
		board  [3][3]int
		winner int
	}{
		{
			"winner is blue",
			[3][3]int{
				{blue, red, empty},
				{empty, blue, red},
				{red, empty, blue},
			},
			blue,
		},
		{
			"winner is red",
			[3][3]int{
				{red, red, red},
				{empty, blue, blue},
				{empty, empty, blue},
			},
			red,
		},
		{
			"there is no winner",
			[3][3]int{
				{red, blue, red},
				{red, blue, blue},
				{blue, red, blue},
			},
			empty,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.winner, HasWonTicTacToe(data.board))
		})
	}
}
