package chapter3

import (
	"errors"
)

type MinAwareStack struct {
	top *StackNode
	min *Stack
}

func (stack *MinAwareStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot pop an empty stack")
	}

	data := stack.top.data

	stack.top = stack.top.next
	stack.min.top = stack.min.top.next

	return data, nil
}

func (stack *MinAwareStack) IsEmpty() bool {
	return stack == nil || *stack == MinAwareStack{} || stack.top == nil
}

func (stack *MinAwareStack) Push(data int) {
	node := &StackNode{data, stack.top}
	if stack.min == nil {
		stack.min = &Stack{node}
	} else if stack.min.top.data < data {
		stack.min.top = &StackNode{stack.min.top.data, stack.min.top}
	} else {
		stack.min.top = &StackNode{data, stack.min.top}
	}

	stack.top = node
}

func (stack *MinAwareStack) Min() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot get the min of a an empty stack")
	}

	return stack.min.top.data, nil
}

func (stack *MinAwareStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot peak an empty stack")
	}

	return stack.top.data, nil
}
