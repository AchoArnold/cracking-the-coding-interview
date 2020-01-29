package chapter3

import (
	"errors"
)

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	first, last *QueueNode
}

func (queue *Queue) Add(data int) {
	node := &QueueNode{data, nil}

	if queue.last != nil {
		queue.last.next = node
	}

	queue.last = node

	if queue.first == nil {
		queue.first = queue.last
	}
}

func (queue *Queue) IsEmpty() bool {
	return queue == &Queue{} || (queue.first == nil && queue.last == nil)
}

func (queue *Queue) Remove() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("cannot remove from an empty stack")
	}

	data := queue.first.data
	queue.first = queue.first.next

	if queue.first == nil {
		queue.last = nil
	}

	return data, nil
}

func (queue *Queue) peek() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("cannot peak an empty stack")
	}

	return queue.first.data, nil
}
