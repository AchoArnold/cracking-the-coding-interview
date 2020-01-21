package chapter2

import (
	"fmt"
)

type SinglyLinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	data int
}

func (linkedList *SinglyLinkedList) appendToTail(data int) {
	end := Node{nil, data}

	currentNode := linkedList.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &end
}

func (linkedList *SinglyLinkedList) appendToHead(data int) {
	head := Node{linkedList.head, data}

	linkedList.head = &head
}

func (linkedList *SinglyLinkedList) deleteNode(data int) {
	node := linkedList.head

	if node.data == data {
		*linkedList.head = *node.next
	}

	for node.next != nil {
		if node.next.data == data {
			*node.next = *node.next.next
			return
		}

		node = node.next
	}
}

func (linkedList *SinglyLinkedList) print() string {
	node := linkedList.head

	outputString := fmt.Sprintf("%d->", node.data)
	for node.next != nil {
		outputString += fmt.Sprintf("%d->", node.next.data)
		node = node.next
	}

	outputString += fmt.Sprintf("nil\n")

	return outputString
}
