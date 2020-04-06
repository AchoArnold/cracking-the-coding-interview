package chapter10

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testArray := makeTestArray(10000)
	MergeSort(testArray)
	assertIsSorted(t, testArray)
}

func makeTestArray(size int) []int {
	testArray := make([]int, size)

	for i := 0; i < size; i++ {
		testArray[i] = rand.Int() % 9973
	}

	return testArray
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testArray := makeTestArray(10)
		MergeSort(testArray)
	}
}

func assertIsSorted(t *testing.T, array []int) {
	start := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < start {
			t.Errorf("%d is less than %d with index %d and %d", array[i], start, i, i-1)
		}

		start = array[i]
	}
}
