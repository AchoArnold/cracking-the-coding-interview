package chapter4

import "math"

func (node *ParentAwareBinaryTreeNode) CommonAncestor(node1Value interface{}, node2Value interface{}) *ParentAwareBinaryTreeNode {
	node1 := node.getNodeFromInterface(node1Value)
	node2 := node.getNodeFromInterface(node2Value)

	delta := Depth(node1) - Depth(node2)

	first := node1
	if delta > 0 {
		first = node2
	}

	second := node2
	if delta > 0 {
		second = node1
	}

	second = GoUpBy(second, int(math.Abs(float64(delta))))

	for first != second && first != nil && second != nil {
		first = first.Parent
		second = second.Parent
	}

	if first == nil || second == nil {
		return nil
	}

	return first
}

func GoUpBy(node *ParentAwareBinaryTreeNode, delta int) *ParentAwareBinaryTreeNode {
	for delta > 0 && node != nil {
		node = node.Parent
		delta--
	}

	return node
}

func Depth(node *ParentAwareBinaryTreeNode) int {
	depth := 0

	for node != nil {
		node = node.Parent
		depth++
	}

	return depth
}

func (node *ParentAwareBinaryTreeNode) getNodeFromInterface(value interface{}) *ParentAwareBinaryTreeNode {
	var nodes = node.GetNodesAsSlice()

	for _, currentNode := range nodes {
		if currentNode.Value == value {
			return currentNode
		}
	}

	return nil
}

func (node *ParentAwareBinaryTreeNode) GetNodesAsSlice() []*ParentAwareBinaryTreeNode {
	if node == nil {
		return []*ParentAwareBinaryTreeNode{}
	}
	return append(node.Left.GetNodesAsSlice(), append([]*ParentAwareBinaryTreeNode{node}, node.Right.GetNodesAsSlice()...)...)
}
