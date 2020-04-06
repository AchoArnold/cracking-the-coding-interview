package chapter10

import "testing"

func TestMergeSorted(t *testing.T) {
	a := makeTestArray(10)
	b := makeTestArray(10)

	MergeSort(a)
	MergeSort(b)

	a = append(a, make([]int, 10)...)

	MergeSorted(a, b, 10, 10)

	assertIsSorted(t, a)
}
