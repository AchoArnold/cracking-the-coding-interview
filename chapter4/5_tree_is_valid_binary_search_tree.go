package chapter4

import "math"

func (root *BinaryTreeNode) IsBinarySearchTree() bool {
	return root.IsBinarySearchTreeWithMinMax(math.MinInt32, math.MaxInt32)
}

func (root *BinaryTreeNode) IsBinarySearchTreeWithMinMax(min int, max int) bool {
	if root == nil {
		return true
	}

	if (min != math.MinInt32 && root.Value.(int) <= min) || (max != math.MaxInt32 && root.Value.(int) > max) {
		return false
	}

	if (!root.Left.IsBinarySearchTreeWithMinMax(min, root.Value.(int))) || !root.Right.IsBinarySearchTreeWithMinMax(root.Value.(int), max) {
		return false
	}

	return true
}
