package chapter4

import (
	"math"
)

func GetBinaryTreeFromSortedArray(array []int) *BinaryTreeNode {
	return createMinimalBinarySubTreeFromArray(IntToInterfaceSlice(array), 0, len(array)-1)
}

func createMinimalBinarySubTreeFromArray(array []interface{}, start int, end int) *BinaryTreeNode {
	if end < start {
		return nil
	}

	var mid = int(math.Ceil(float64(end+start) / float64(2)))

	return &BinaryTreeNode{
		Value: array[mid],
		Left:  createMinimalBinarySubTreeFromArray(array, start, mid-1),
		Right: createMinimalBinarySubTreeFromArray(array, mid+1, end),
	}
}
