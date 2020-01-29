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

func (linkedList *SinglyLinkedList) deleteNode(data int) *SinglyLinkedList {
	node := linkedList.head

	if node.data == data {
		*linkedList.head = *node.next
	}

	for node.next != nil {
		if node.next.data == data {
			*node.next = *node.next.next
			return linkedList
		}

		node = node.next
	}

	return linkedList
}

func (linkedList *SinglyLinkedList) String() string {
	node := linkedList.head

	outputString := fmt.Sprintf("%d->", node.data)
	for node.next != nil {
		outputString += fmt.Sprintf("%d->", node.next.data)
		node = node.next
	}

	outputString += fmt.Sprintf("nil")

	return outputString
}

func makeListForRange(upperBound int) *SinglyLinkedList {
	list := SinglyLinkedList{&Node{nil, 0}}
	for index := 1; index < upperBound; index++ {
		list.appendToTail(index)
	}

	return &list
}

func makeFromSlice(values []int) *SinglyLinkedList {
	list := SinglyLinkedList{&Node{nil, values[0]}}
	for index := 1; index < len(values); index++ {
		list.appendToTail(values[index])
	}

	return &list
}

func (linkedList *SinglyLinkedList) Count() int {
	if (*linkedList == SinglyLinkedList{}) {
		return 0
	}

	node := linkedList.head
	count := 0

	for node != nil {
		count++
		node = node.next
	}

	return count
}

func (linkedList *SinglyLinkedList) IsEmpty() bool {
	return *linkedList == SinglyLinkedList{}
}

func (linkedList *SinglyLinkedList) Clone() *SinglyLinkedList {
	newList := &SinglyLinkedList{}
	if linkedList.IsEmpty() {
		return newList
	}

	node := linkedList.head
	newList.head = &Node{nil, node.data}

	for node.next != nil {
		newList.appendToTail(node.next.data)
		node = node.next
	}

	return newList
}
func (linkedList *SinglyLinkedList) Reverse() *SinglyLinkedList {
	node := linkedList.head
	newList := SinglyLinkedList{&Node{nil, node.data}}

	for node.next != nil {
		newList.appendToHead(node.next.data)
		node = node.next
	}

	*linkedList = newList

	return linkedList
}

func (linkedList *SinglyLinkedList) Slice() []int {
	var data []int
	node := linkedList.head

	for node != nil {
		data = append(data, node.data)
		node = node.next
	}

	return data
}
