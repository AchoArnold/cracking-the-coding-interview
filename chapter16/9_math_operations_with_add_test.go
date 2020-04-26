package chapter16

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyWithOnlyAddition(t *testing.T) {
	testData := []struct {
		name           string
		x, y, expected int
	}{
		{"zero", 1, 0, 0},
		{"same", 2, 2, 4},
		{"mixed", -3, 2, -6},
		{"all negative", -2, -5, 10},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, MultiplyWithOnlyAddition(data.x, data.y))
		})
	}
}

func TestSubtractWithOnlyAddition(t *testing.T) {
	testData := []struct {
		name           string
		x, y, expected int
	}{
		{"zero", 0, 0, 0},
		{"same", 2, 2, 0},
		{"mixed", -3, 2, -5},
		{"all negative", -2, -5, 3},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, SubtractWithOnlyAddition(data.x, data.y))
		})
	}
}

func TestDivideWithOnlyAddition(t *testing.T) {
	testData := []struct {
		name           string
		x, y, expected int
	}{
		{"zero", 0, 0, math.MaxInt32},
		{"number and zero", 332, 0, math.MaxInt32},
		{"same", 2, 2, 1},
		{"mixed less", -3, 2, -1},
		{"all negative greater", -2, -5, 0},
		{"all negative equal", -2, -2, 1},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, DivideWithOnlyAddition(data.x, data.y))
		})
	}
}

func TestFlipSignWithAddition(t *testing.T) {
	testData := []struct {
		name        string
		x, expected int
	}{
		{"zero", 0, 0},
		{"positive", 2, -2},
		{"negative", -3, 3},
		{"one", 1, -1},
		{"negative one", -1, 1},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, FlipSignWithAddition(data.x))
		})
	}
}
