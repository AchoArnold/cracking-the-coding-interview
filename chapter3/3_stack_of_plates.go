package chapter3

import (
	"errors"
	"fmt"
)

const limit = 3

type MultiDimensionalStack struct {
	top *NodeStack
}

type NodeStack struct {
	data *LengthAwareStack
	next *NodeStack
}
type LengthAwareStack struct {
	top    *StackNode
	length int
}

func (stack *MultiDimensionalStack) Push(data int) {
	if stack.top == nil {
		stack.top = &NodeStack{
			data: &LengthAwareStack{&StackNode{data, nil}, 1},
			next: nil,
		}
	} else {
		currentStack := stack.top
		if currentStack.data.length == limit {
			stack.top = &NodeStack{
				data: &LengthAwareStack{&StackNode{data, nil}, 1},
				next: stack.top,
			}
		} else {
			currentStack.data.length += 1
			currentStack.data.top = &StackNode{data, currentStack.data.top}
		}
	}
}

func (stack *MultiDimensionalStack) IsEmpty() bool {
	return stack == nil || *stack == MultiDimensionalStack{} || stack.top == nil
}

func (stack *MultiDimensionalStack) PopAt(index int) (int, error) {
	topNode := stack.top
	counter := 0

	for topNode != nil {
		if counter == index {
			val := topNode.data.top.data
			topNode.data.top = topNode.data.top.next
			return val, nil
		}
		counter++
		topNode = topNode.next
	}
	return -1, errors.New(fmt.Sprintf("invalid index %d", index))
}

func (stack *MultiDimensionalStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot pop from an empty stack")
	}

	currentStack := stack.top
	val := currentStack.data.top.data
	currentStack.data.top = currentStack.data.top.next
	currentStack.data.length -= 1

	if currentStack.data.length == 0 {
		stack.top = stack.top.next
	}

	return val, nil
}

func (stack *MultiDimensionalStack) Print() string {
	if stack.IsEmpty() {
		return "nil"
	}

	result := ""
	topNode := stack.top

	for topNode != nil {
		node := topNode.data.top
		for node != nil {
			result += fmt.Sprintf("%d->", node.data)
			node = node.next
		}
		result += fmt.Sprintf("nil\n")
		topNode = topNode.next
	}

	return result
}

func (stack *MultiDimensionalStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("cannot peak an empty stack")
	}

	return stack.top.data.top.data, nil
}
