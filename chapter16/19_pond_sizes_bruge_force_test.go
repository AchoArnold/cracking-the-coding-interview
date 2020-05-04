package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pondSizesTestData = []struct {
	name  string
	land  [][]int
	sizes []int
}{
	{
		"valid land",
		[][]int{
			{0, 2, 1, 0},
			{0, 1, 0, 1},
			{1, 1, 0, 1},
			{0, 1, 0, 1},
		},
		[]int{2, 4, 1},
	},
	{
		"empty",
		[][]int{},
		[]int{},
	},
}

func TestPondSizesBruteForce(t *testing.T) {
	for _, data := range pondSizesTestData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.sizes, PondSizesBruteForce(data.land))
		})
	}
}

func TestPondSizesOptimized(t *testing.T) {
	for _, data := range pondSizesTestData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.sizes, PondSizesOptimized(data.land))
		})
	}
}

func BenchmarkPondSizesBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PondSizesBruteForce(pondSizesTestData[0].land)
	}
}

func BenchmarkPondSizesOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PondSizesOptimized(pondSizesTestData[0].land)
	}
}
