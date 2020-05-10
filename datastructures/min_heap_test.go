package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinHeap(t *testing.T) {
	assertMinHeapEquals(t, &MinHeap{[]int{}}, NewMinHeap())
	assertIsMinHeap(t, NewMinHeap())
}

func TestBuildMinHeapFromArray(t *testing.T) {
	testData := []struct {
		name  string
		input []int
	}{
		{
			"empty array",
			[]int{},
		},
		{
			"array with elements",
			[]int{2, 4, 8, 14, 6, 7, 1, 9, 10, 3},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assertIsMinHeap(t, BuildMinHeapFromArray(data.input))
		})
	}
}

func TestMinHeap_Size(t *testing.T) {
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
			assert.Equal(t, data.expectedSize, BuildMinHeapFromArray(data.input).Size())
		})
	}
}

func TestMinHeap_Insert(t *testing.T) {
	testData := []struct {
		name        string
		inputHeap   *MinHeap
		insertInput []int
	}{
		{
			"empty heap insert 1",
			NewMinHeap(),
			[]int{3},
		},
		{
			"Add to already existing heap",
			&MinHeap{[]int{1, 3, 2, 9}},
			[]int{14, 7},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			for _, val := range data.insertInput {
				data.inputHeap.Insert(val)
			}

			assertIsMinHeap(t, data.inputHeap)
		})
	}
}

func TestMinHeap_Pop(t *testing.T) {
	t.Run("it returns and error when the heap is empty", func(t *testing.T) {
		heap := NewMinHeap()
		_, err := heap.Pop()
		assert.Error(t, err)
	})

	t.Run("it removes the last element from the heap", func(t *testing.T) {
		elements := []int{1, 3, 2, 9, 4, 8, 7, 14, 10, 6}
		heap := &MinHeap{elements}
		val, err := heap.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 1, val)
		assert.Equal(t, len(elements)-1, heap.Size())
	})
}

func TestMinHeap_Min(t *testing.T) {
	t.Run("it returns and error when the heap is empty", func(t *testing.T) {
		heap := NewMinHeap()
		_, err := heap.Min()
		assert.Error(t, err)
	})

	t.Run("it returns the last element from the heap", func(t *testing.T) {
		elements := []int{1, 3, 2, 9, 4, 8, 7, 14, 10, 6}
		heap := &MinHeap{elements}
		val, err := heap.Min()

		assert.NoError(t, err)
		assert.Equal(t, 1, val)
		assert.Equal(t, len(elements), heap.Size())
	})
}

func assertMinHeapEquals(t *testing.T, expected *MinHeap, actual *MinHeap) {
	assert.Equal(t, expected.elements, actual.elements)
}

func assertIsMinHeap(t *testing.T, heap *MinHeap) {
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
				if elements[left] < elements[parent] {
					t.Errorf("parent %d is greater than child %d", elements[parent], elements[left])
				}

				newQueue = append(newQueue, left)
			}

			right := parent*2 + 2
			if right < len(elements) {
				if elements[right] < elements[parent] {
					t.Errorf("parent %d is greater than child %d", elements[parent], elements[right])
				}
				newQueue = append(newQueue, right)
			}
		}
		parentsQueue = newQueue
	}
}
