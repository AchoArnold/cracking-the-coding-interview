package chapter2

func (linkedList *SinglyLinkedList) removeDuplicates() *SinglyLinkedList {
	node := linkedList.head

	for node.next != nil {
		currentHead := *node
		currentData := node.data

		for node != nil && node.next != nil {
			if node.next.data == currentData {
				if node.next.next != nil {
					*node.next = *node.next.next
				} else {
					node.next = nil
				}
			}

			node = node.next
		}

		node = currentHead.next
	}

	return linkedList
}
