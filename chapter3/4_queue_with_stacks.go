package chapter3

import (
	"errors"
)

type QueueUsingStacks struct {
	new *Stack
	old *Stack
}


func (queue *QueueUsingStacks) Add(data int) {
	if queue.IsEmpty() {
		queue.new = &Stack{nil}
	}
	queue.new.Push(data)
}

func (queue *QueueUsingStacks) IsEmpty() bool {
	return queue == nil || queue.new == nil || *queue == QueueUsingStacks{}
}

func (queue *QueueUsingStacks) Remove() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("cannot remove from an empty queue")
	}

	queue.old = &Stack{nil}
	for  {
		val, err := queue.new.Pop()
		if err != nil  {
			break
		}
		queue.old.Push(val)
	}

	val,_ := queue.old.Pop()

	queue.new = &Stack{nil}
	for  {
		val, err := queue.old.Pop()
		if err != nil {
			break
		}

		queue.new.Push(val)
	}

	return val, nil
}


func (queue *QueueUsingStacks) Peek() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("cannot peek from an empty queue")
	}

	queue.old = &Stack{nil}
	for  {
		val, err := queue.new.Pop()
		if err != nil  {
			break
		}
		queue.old.Push(val)
	}

	val,_ := queue.old.Peek()

	queue.new = &Stack{nil}
	for  {
		val, err := queue.old.Pop()
		if err != nil {
			break
		}

		queue.new.Push(val)
	}

	return val, nil
}
