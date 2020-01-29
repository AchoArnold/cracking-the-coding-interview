package chapter2

func (linkedList *SinglyLinkedList) GetLoopingNode() *Node {
	seenNodes := map[*Node]bool{}

	node := linkedList.head

	for node != nil {
		if _, ok := seenNodes[node]; !ok {
			seenNodes[node] = true
		} else {
			return node
		}

		node = node.next
	}

	return nil
}
