package chapter4

import "container/list"

func CreateLevelLinkedList(root *BinaryTreeNode) []*list.List {
	var result []*list.List

	current := list.New()

	if root != nil {
		current.PushBack(root)
	}

	for current.Len() > 0 {
		result = append(result, current)

		parents := current
		current = list.New()

		for parent := parents.Front(); parent != nil; parent = parent.Next() {
			node := parent.Value.(*BinaryTreeNode)

			if node.Left != nil {
				current.PushBack(node.Left)
			}

			if node.Right != nil {
				current.PushBack(node.Right)
			}
		}
	}

	return result
}
