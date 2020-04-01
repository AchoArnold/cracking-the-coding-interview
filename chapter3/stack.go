package chapter3

import (
	"errors"
)

type StackNode struct {
	data int
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot pop an empty stack")
	}

	data := stack.top.data
	stack.top = stack.top.next

	return data, nil
}

func (stack *Stack) IsEmpty() bool {
	return stack == &Stack{} || stack.top == nil
}

func (stack *Stack) Push(data int) {
	node := &StackNode{data, stack.top}
	stack.top = node
}

func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot peak an empty stack")
	}

	return stack.top.data, nil
}

func (stack *Stack) Empty() {
	for {
		_, err := stack.Pop()
		if err != nil {
			break
		}
	}
}

func (stack *Stack) ToArray() []int {
	var array []int
	if stack.IsEmpty() {
		return array
	}

	node := stack.top

	for node != nil {
		array = append(array, node.data)
		node = node.next
	}

	return array
}
