package chapter4

func GetBinaryTreeFromSortedArray(array []int) *BinaryTreeNode {
	return createMinimalBinarySubTreeFromArray(IntToInterfaceSlice(array), 0, len(array) -1)
}


func createMinimalBinarySubTreeFromArray(array []interface{}, start int, end int) *BinaryTreeNode {
	if end <  start {
		return nil
	}

	var mid = (start + end) /2

	var node = &BinaryTreeNode{
		Value: array[mid],
	}

	node.Left = createMinimalBinarySubTreeFromArray(array, start, mid-1)
	node.Right = createMinimalBinarySubTreeFromArray(array, mid+1, end)

	return node
}