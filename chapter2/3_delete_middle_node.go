package chapter2

func (linkedList *SinglyLinkedList) deleteMiddleNode(node *Node) *SinglyLinkedList{
	*node = *node.next
	return linkedList
}