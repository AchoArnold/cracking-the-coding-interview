package chapter17

// BitNode is a node with 2 links
type BitNode struct {
	value int
	node1 *BitNode
	node2 *BitNode
}

// ConvertTreeToDoublyLinkedList converts a list to a circular linked list and then breaks the circular connection
func ConvertTreeToDoublyLinkedList(node *BitNode) *BitNode {
	head := convertToCircular(node)
	if head != nil {
		head.node1.node2 = nil
		head.node1 = nil
	}

	return head
}

func convertToCircular(root *BitNode) *BitNode {
	if root == nil {
		return nil
	}

	part1 := convertToCircular(root.node1)
	part3 := convertToCircular(root.node2)

	if part1 == nil && part3 == nil {
		root.node1 = root
		root.node2 = root
		return root
	}

	var tail3 *BitNode
	if part3 != nil {
		tail3 = part3.node1
	}

	// join left to root
	if part1 == nil {
		bitNodeConcat(part3.node1, root)
	} else {
		bitNodeConcat(part1.node1, root)
	}

	// join right to root
	if part3 == nil {
		bitNodeConcat(root, part1)
	} else {
		bitNodeConcat(root, part3)
	}

	// join right to left
	if part1 != nil && part3 != nil {
		bitNodeConcat(tail3, part1)
	}

	if part1 == nil {
		return root
	}

	return part1
}

func bitNodeConcat(x, y *BitNode) {
	x.node2 = y
	y.node1 = x
}
