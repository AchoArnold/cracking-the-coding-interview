package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	assertHeapEquals(t, &MaxHeap{[]int{}}, NewMaxHeap())
	assertIsMaxHeap(t, NewMaxHeap())
}

func TestBuildMaxHeapFromArray(t *testing.T) {
	testData := []struct {
		name         string
		input        []int
		expectedHeap *MaxHeap
	}{
		{
			"empty array",
			[]int{},
			&MaxHeap{[]int{}},
		},
		{
			"array with elements",
			[]int{2, 4, 8, 14, 6, 7, 1, 9, 10, 3},
			&MaxHeap{
				[]int{14, 10, 7, 9, 6, 4, 1, 2, 8, 3},
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assertHeapEquals(t, data.expectedHeap, BuildMaxHeapFromArray(data.input))
		})
	}
}

func TestMaxHeap_Size(t *testing.T) {
	testData := []struct {
		name         string
		input        []int
		expectedSize int
	}{
		{
			"empty array",
			[]int{},
			0,
		},
		{
			"array with elements",
			[]int{2, 4, 8, 14, 6, 7, 1, 9, 10, 3},
			10,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedSize, BuildMaxHeapFromArray(data.input).Size())
		})
	}
}

func TestMaxHeap_Insert(t *testing.T) {
	testData := []struct {
		name        string
		inputHeap   *MaxHeap
		insertInput []int
	}{
		{
			"empty heap insert 1",
			NewMaxHeap(),
			[]int{3},
		},
		{
			"Add to already existing heap",
			&MaxHeap{[]int{8, 2, 4}},
			[]int{14, 7},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			for _, val := range data.insertInput {
				data.inputHeap.Insert(val)
			}

			assertIsMaxHeap(t, data.inputHeap)
		})
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	t.Run("it returns and error when the heap is empty", func(t *testing.T) {
		heap := NewMaxHeap()
		_, err := heap.Pop()
		assert.Error(t, err)
	})

	t.Run("it removes the last element from the heap", func(t *testing.T) {
		elements := []int{14, 10, 7, 9, 6, 4, 1, 2, 8, 3}
		heap := &MaxHeap{elements}
		val, err := heap.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 14, val)
		assert.Equal(t, len(elements)-1, heap.Size())
	})
}

func TestMaxHeap_Max(t *testing.T) {
	t.Run("it returns and error when the heap is empty", func(t *testing.T) {
		heap := NewMaxHeap()
		_, err := heap.Max()
		assert.Error(t, err)
	})

	t.Run("it returns the last element from the heap", func(t *testing.T) {
		elements := []int{14, 10, 7, 9, 6, 4, 1, 2, 8, 3}
		heap := &MaxHeap{elements}
		val, err := heap.Max()

		assert.NoError(t, err)
		assert.Equal(t, 14, val)
		assert.Equal(t, len(elements), heap.Size())
	})
}

func assertHeapEquals(t *testing.T, expected *MaxHeap, actual *MaxHeap) {
	assert.Equal(t, expected.elements, actual.elements)
}

func assertIsMaxHeap(t *testing.T, heap *MaxHeap) {
	elements := heap.elements

	if len(elements) == 0 {
		return
	}

	parentsQueue := []int{0}
	for len(parentsQueue) != 0 {
		var newQueue []int
		for _, parent := range parentsQueue {
			left := parent*2 + 1
			if left < len(elements) {
				if elements[left] > elements[parent] {
					t.Errorf("parent %d is less than child %d", elements[parent], elements[left])
				}

				newQueue = append(newQueue, left)
			}

			right := parent*2 + 2
			if right < len(elements) {
				if elements[right] > elements[parent] {
					t.Errorf("parent %d is less than child %d", elements[parent], elements[right])
				}
				newQueue = append(newQueue, right)
			}
		}
		parentsQueue = newQueue
	}
}
