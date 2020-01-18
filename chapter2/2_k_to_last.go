package chapter2

func (linkedList *SinglyLinkedList) kToLast(k int) int {
	node := linkedList.head

	data := []int{node.data}

	for node.next != nil  {
		data = append(data, node.next.data)
		node = node.next
	}

	return data[len(data) - k -1]
}