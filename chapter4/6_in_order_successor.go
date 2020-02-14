package chapter4

import "math"

type ParentAwareBinaryTreeNode struct {
	Value  interface{}
	Left   *ParentAwareBinaryTreeNode
	Right  *ParentAwareBinaryTreeNode
	Parent *ParentAwareBinaryTreeNode
}

func createParentAwareBinarySearchTreeFromArray(parent *ParentAwareBinaryTreeNode, array []interface{}, start int, end int) *ParentAwareBinaryTreeNode {
	if end < start {
		return nil
	}

	var mid = int(math.Ceil(float64(end+start) / float64(2)))

	node := &ParentAwareBinaryTreeNode{
		Value:  array[mid],
		Parent: parent,
	}

	node.Left = createParentAwareBinarySearchTreeFromArray(node, array, start, mid-1)
	node.Right = createParentAwareBinarySearchTreeFromArray(node, array, mid+1, end)

	return node
}

func CreatParentAwareBinarySearchTree(array []interface{}) *ParentAwareBinaryTreeNode {
	return createParentAwareBinarySearchTreeFromArray(nil, array, 0, len(array)-1)
}

func (node *ParentAwareBinaryTreeNode) InOrderSuccessor() *ParentAwareBinaryTreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return node.Right.leftMostChild()
	}

	q := node
	x := node.Parent

	for x != nil && x.Left != q {
		q = x
		x = x.Parent
	}

	return x
}

func (node *ParentAwareBinaryTreeNode) leftMostChild() *ParentAwareBinaryTreeNode {
	if node == nil {
		return nil
	}

	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}
