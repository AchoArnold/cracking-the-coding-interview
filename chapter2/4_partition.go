package chapter2

func (linkedList *SinglyLinkedList) partition(partition int) {
	var lessThanPartitionTail, greaterThanPartitionHead, lessThanPartitionHead *Node

	head := linkedList.head

	for head != nil {
		node := *head
		if node.data < partition {
			if lessThanPartitionTail == nil {
				lessThanPartitionTail = &node
				lessThanPartitionHead = &node
			} else {
				lessThanPartitionTail.next = &node
				lessThanPartitionTail = lessThanPartitionTail.next
			}
		} else {
			node.next = nil
			if greaterThanPartitionHead == nil {
				greaterThanPartitionHead = &node
			} else {
				node.next = greaterThanPartitionHead
				greaterThanPartitionHead = &node
			}
		}

		head = head.next
	}

	if lessThanPartitionTail != nil && lessThanPartitionHead != nil {
		lessThanPartitionTail.next = greaterThanPartitionHead
		*linkedList.head = *lessThanPartitionHead
	} else if greaterThanPartitionHead != nil {
		*linkedList.head = *greaterThanPartitionHead
	}
}
