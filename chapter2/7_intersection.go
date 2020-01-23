package chapter2

func intersection(list1, list2 *SinglyLinkedList) *Node {
	node1 := list1.head
	node2 := list2.head

	seenNodes := map[*Node]int{}

	for node1 != nil || node2 != nil {

		if value, exists := seenNodes[node1]; exists && value != 1 {
			return node1
		} else {
			seenNodes[node1] = 1
		}

		if value, exists := seenNodes[node2]; exists && value != 2 {
			return node2
		} else {
			seenNodes[node2] = 2
		}

		if node1 != nil {
			node1 = node1.next
		}
		if node2 != nil {
			node2 = node2.next
		}
	}

	return nil
}
