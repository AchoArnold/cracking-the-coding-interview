package chapter4

import (
	"math"
)

func (root *BinaryTreeNode) GetBinaryHeight() int {
	if root == nil {
		return -1
	}

	return int(math.Max(float64(root.Left.GetBinaryHeight()), float64(root.Right.GetBinaryHeight()))) + 1
}

func (root *BinaryTreeNode) IsBalanced() bool {
	if root == nil {
		return true
	}

	heightDiff := root.Left.GetBinaryHeight() - root.Right.GetBinaryHeight()

	if math.Abs(float64(heightDiff)) > 1 {
		return false
	}

	return root.Left.IsBalanced() && root.Right.IsBalanced()
}

func (root *BinaryTreeNode) CheckBalancedHeight() int {
	if root == nil {
		return -1
	}

	leftHeight := root.Left.CheckBalancedHeight()
	if leftHeight == math.MinInt32 {
		return math.MinInt32
	}

	rightHeight := root.Right.CheckBalancedHeight()
	if rightHeight == math.MinInt32 {
		return math.MinInt32
	}

	heightDiff := leftHeight - rightHeight
	if math.Abs(float64(heightDiff)) > 1 {
		return math.MinInt32
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func (root *BinaryTreeNode) IsBalancedUsingHeightCheck() bool {
	return root.CheckBalancedHeight() != math.MinInt32
}
