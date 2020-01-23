package chapter2

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"

)

func TestSlice(t *testing.T) {
	list := makeListForRange(4)

	assert.Equal(t, []int{0, 1, 2, 3}, list.Slice())
}

func TestReverse(t *testing.T) {
	list := makeListForRange(4)

	assert.Equal(t, []int{3, 2, 1, 0}, list.Reverse().Slice())
}

func TestString(t *testing.T) {
	list := makeListForRange(4)
	assert.Equal(t, "0->1->2->3->nil", list.String())
}

func TestMakeFromSlice(t *testing.T) {
	slices := [][]int{
		[]int{1, 2, 3, 5},
		[]int{1, 2, 4},
	}

	for _, slice := range slices {
		t.Run(strconv.Itoa(len(slice)) + " items slice", func(t *testing.T) {
			assert.Equal(t, makeFromSlice(slice).Slice(), slice)
		})
	}
}

func TestCount(t *testing.T) {
	slices := [][]int{
		[]int{1, 2, 3, 5},
		[]int{1, 2, 4},
	}

	for _, slice := range slices {
		t.Run(strconv.Itoa(len(slice))+" items slice", func(t *testing.T) {
			assert.Equal(t, makeFromSlice(slice).Count(), len(slice))
		})
	}

	t.Run("Testing Empty List", func(t *testing.T) {
		var list SinglyLinkedList
		assert.Equal(t, list.Count(), 0)
	})
}

func TestDeleteNode(t *testing.T) {
	testData := []struct{
		name string
		listData []int
		dataToDelete int
		expectedList []int
 	} {
 		{
 			"Deleting head",
 			[]int{1,2,3},
 			1,
 			[]int{2,3},
		},
		{
			"Deleting middle",
			[]int{1,2,3},
			2,
			[]int{1,3},
		},
		{
			"Deleting tail",
			[]int{1,2},
			3,
			[]int{1,2},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expectedList, makeFromSlice(data.listData).deleteNode(data.dataToDelete).Slice())
		})
	}
}