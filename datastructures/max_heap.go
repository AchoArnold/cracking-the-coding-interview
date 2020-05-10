package datastructures

import (
	"math"

	"github.com/go-errors/errors"
)

// MaxHeap represents a max heap
type MaxHeap struct {
	elements []int
}

// NewMaxHeap creates a new max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{[]int{}}
}

// BuildMaxHeapFromArray creates a new max heap from an array of values
func BuildMaxHeapFromArray(input []int) *MaxHeap {
	heap := NewMaxHeap()

	for _, val := range input {
		heap.Insert(val)
	}

	return heap
}

// Size represents the current size of the heap
func (heap MaxHeap) Size() int {
	return len(heap.elements)
}

// Insert inserts a value int the heap
func (heap *MaxHeap) Insert(val int) {
	heap.elements = append(heap.elements, val)
	heap.maxHeapify(heap.getParent(heap.lastIndex()))
}

func (heap *MaxHeap) isRoot(i int) bool {
	return i == 0
}

func (heap *MaxHeap) lastIndex() int {
	return heap.Size() - 1
}

func (heap *MaxHeap) getParent(i int) int {
	return (i - 1) / 2
}

// Pop removes the least element in the heap
func (heap *MaxHeap) Pop() (int, error) {
	if heap.Size() == 0 {
		return 0, errors.New("heap underflow")
	}

	max := heap.elements[0]
	heap.elements[0] = heap.elements[heap.Size()-1]
	heap.elements = heap.elements[:heap.Size()-1]
	heap.maxHeapify(0)

	return max, nil
}

// Max returns the max element in the heap
func (heap *MaxHeap) Max() (int, error) {
	if heap.Size() == 0 {
		return math.MaxInt32, errors.New("no element in heap")
	}

	return heap.elements[0], nil
}

func (heap *MaxHeap) maxHeapify(i int) {
	leftIndex := (i * 2) + 1
	rightIndex := (i * 2) + 2

	largestIndex := i
	if leftIndex < heap.Size() && heap.elements[leftIndex] > heap.elements[i] {
		largestIndex = leftIndex
	}

	if rightIndex < heap.Size() && heap.elements[rightIndex] > heap.elements[largestIndex] {
		largestIndex = rightIndex
	}

	if largestIndex != i {
		heap.swapIndex(i, largestIndex)
		if !heap.isRoot(i) {
			if heap.elements[heap.getParent(i)] < heap.elements[i] {
				heap.maxHeapify(heap.getParent(i))
			}
		}
	}
}

func (heap *MaxHeap) swapIndex(i, j int) {
	temp := heap.elements[j]
	heap.elements[j] = heap.elements[i]
	heap.elements[i] = temp
}
