package chapter4

import (
	"math/rand"
)

type RandomTreeNode struct {
	Value       int
	Left, Right *RandomTreeNode
	Size        int
}

func NewRandomTreeNode(value int) *RandomTreeNode {
	return &RandomTreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
		Size:  1,
	}
}

func (root *RandomTreeNode) GetRandomNode() *RandomTreeNode {
	leftSize := 0
	if root.Left != nil {
		leftSize = root.Left.Size
	}

	index := 0
	if root.Size > 0 {
		index = rand.Intn(root.Size)
	}

	if index < leftSize {
		return root.Left.GetRandomNode()
	}

	if index == leftSize {
		return root
	}

	return root.Right.GetRandomNode()
}

func (root *RandomTreeNode) InsertInOrder(value int) {
	if value < root.Value {
		if root.Left == nil {
			root.Left = NewRandomTreeNode(value)
		} else {
			root.Left.InsertInOrder(value)
		}
	} else {
		if root.Right == nil {
			root.Right = NewRandomTreeNode(value)
		} else {
			root.Right.InsertInOrder(value)
		}
	}

	root.Size++
}

func (root *RandomTreeNode) Find(value int) *RandomTreeNode {
	if value == root.Value {
		return root
	}

	if value < root.Value {
		if root.Left != nil {
			return root.Left.Find(value)
		} else {
			return nil
		}
	}

	if root.Right != nil {
		return root.Right.Find(value)
	}
	return nil
}
