package datastructures

import (
	"math"

	"github.com/go-errors/errors"
)

// MinHeap is the data structure for a min heap
type MinHeap struct {
	elements []int
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{[]int{}}
}

// BuildMinHeapFromArray creates a new min heap from a slice of integers
func BuildMinHeapFromArray(input []int) *MinHeap {
	heap := NewMinHeap()

	for _, val := range input {
		heap.Insert(val)
	}

	return heap
}

// Size returns the size of a mean heap
func (heap MinHeap) Size() int {
	return len(heap.elements)
}

// Insert inserts a value in the min heap
func (heap *MinHeap) Insert(val int) {
	heap.elements = append(heap.elements, val)
	heap.minHeapify(heap.getParent(heap.lastIndex()))
}

func (heap *MinHeap) isRoot(i int) bool {
	return i == 0
}

func (heap *MinHeap) lastIndex() int {
	return heap.Size() - 1
}

func (heap *MinHeap) getParent(i int) int {
	return (i - 1) / 2
}

// Pop returns the min element in the heap and removes it
func (heap *MinHeap) Pop() (int, error) {
	if heap.Size() == 0 {
		return 0, errors.New("heap underflow")
	}

	min := heap.elements[0]
	heap.elements[0] = heap.elements[heap.Size()-1]
	heap.elements = heap.elements[:heap.Size()-1]
	heap.minHeapify(0)

	return min, nil
}

// Min returns the min element in the heap
func (heap *MinHeap) Min() (int, error) {
	if heap.Size() == 0 {
		return math.MinInt32, errors.New("no element in heap")
	}

	return heap.elements[0], nil
}

func (heap *MinHeap) minHeapify(i int) {
	leftIndex := (i * 2) + 1
	rightIndex := (i * 2) + 2

	largestIndex := i
	if leftIndex < heap.Size() && heap.elements[leftIndex] < heap.elements[i] {
		largestIndex = leftIndex
	}

	if rightIndex < heap.Size() && heap.elements[rightIndex] < heap.elements[largestIndex] {
		largestIndex = rightIndex
	}

	if largestIndex != i {
		heap.swapIndex(i, largestIndex)
		if !heap.isRoot(i) {
			if heap.elements[heap.getParent(i)] > heap.elements[i] {
				heap.minHeapify(heap.getParent(i))
			}
		}
	}
}

func (heap *MinHeap) swapIndex(i, j int) {
	temp := heap.elements[j]
	heap.elements[j] = heap.elements[i]
	heap.elements[i] = temp
}
