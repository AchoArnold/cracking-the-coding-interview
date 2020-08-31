package chapter17

import (
	"github.com/daviddengcn/go-villa"
)

// RunningMedianCalculator gets the running median for a stream of items
type RunningMedianCalculator struct {
	minHeap *villa.IntPriorityQueue
	maxHeap *villa.IntPriorityQueue
}

// NewRunningMedianCalculator calculates the running median
func NewRunningMedianCalculator() RunningMedianCalculator {
	comparator := villa.IntCmpFunc(func(a, b int) int {
		if a < b {
			return 1
		}

		if a > b {
			return -1
		}

		return 0
	})

	return RunningMedianCalculator{
		minHeap: villa.NewIntPriorityQueue(villa.IntValueCompare),
		maxHeap: villa.NewIntPriorityQueue(comparator),
	}
}

func (calculator *RunningMedianCalculator) addItem(item int) {
	// Note: addItem maintains a condition that maxHeap.size() >= minHeap.size()
	if calculator.maxHeap.Len() == calculator.minHeap.Len() {
		if calculator.minHeap.Len() > 0 && item > calculator.minHeap.Peek() {
			calculator.maxHeap.Push(calculator.minHeap.Pop())
			calculator.minHeap.Push(item)
		} else {
			calculator.maxHeap.Push(item)
		}
	} else {
		if item < calculator.maxHeap.Peek() {
			calculator.minHeap.Push(calculator.maxHeap.Pop())
			calculator.maxHeap.Push(item)
		} else {
			calculator.minHeap.Push(item)
		}
	}
}

// Median returns the running median
func (calculator *RunningMedianCalculator) Median() float64 {
	// maxHeap is always at least as big as min heap so if maxHeap is empty then meanHeap is also empty
	if calculator.maxHeap.Len() == 0 {
		return 0
	}

	if calculator.maxHeap.Len() == calculator.minHeap.Len() {
		return float64(calculator.minHeap.Peek()+calculator.maxHeap.Peek()) / 2
	}

	// if maxHeap and minHeap are of different sizes, then maxHeap must have one extra element. return that
	return float64(calculator.maxHeap.Peek())
}
