package chapter8

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	t.Run("Testing on a 3x3 maze", func(t *testing.T) {
		maze := [][]int{
			{1, 2, 3},
			{2, 2, 4},
			{1, 3, 3},
		}

		expected := []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}

		assert.True(t, reflect.DeepEqual(expected, GetPath(maze)))
	})

	t.Run("Testing on a nil maze", func(t *testing.T) {
		var maze [][]int
		var expected []Point

		assert.True(t, reflect.DeepEqual(expected, GetPath(maze)))
	})

	t.Run("Testing on a nxm maze", func(t *testing.T) {
		maze := [][]int{
			{1, 2, 3},
			{2, 2, 4},
		}

		expected := []Point{{0, 0}, {1, 0}, {1, 1}, {1, 2}}

		assert.True(t, reflect.DeepEqual(expected, GetPath(maze)))
	})
}
