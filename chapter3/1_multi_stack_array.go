package chapter3

import (
	"errors"
	"fmt"
)

// We assume the 3 stacks of length = 4 each
type MultiArrayStack struct {
	stackIndexes [3]int
	stack        [12]int
}

const index0 = 0
const index1 = 4
const index2 = 8

func NewMultiArrayStack() *MultiArrayStack {
	stack := MultiArrayStack{getDefaultIndexes(), [12]int{}}
	return &stack
}

func getDefaultIndexes() [3]int {
	return [3]int{index0, index1, index2}
}

func (stack *MultiArrayStack) Pop(stackIndex int) (int, error) {
	stack.AssertStackIndexIsWithinRange(stackIndex)

	if stack.IsEmpty(stackIndex) {
		return -1, errors.New(fmt.Sprintf("Cannot pop from stack %d because it is empty", stackIndex))
	}

	value := stack.stack[stack.stackIndexes[stackIndex]-1]
	stack.stackIndexes[stackIndex]--
	return value, nil
}

func (stack *MultiArrayStack) Push(stackIndex int, data int) error {
	stack.AssertStackIndexIsWithinRange(stackIndex)

	if stack.stackIndexes[stackIndex] == getDefaultIndexes()[stackIndex]+4 {
		return errors.New(fmt.Sprintf("Cannot push into stack %d because it is full", stackIndex))
	}

	stack.stack[stack.stackIndexes[stackIndex]] = data
	stack.stackIndexes[stackIndex]++

	return nil
}

func (stack *MultiArrayStack) Peek(stackIndex int, data int) (int, error) {
	stack.AssertStackIndexIsWithinRange(stackIndex)

	if stack.IsEmpty(stackIndex) {
		return -1, errors.New(fmt.Sprintf("Cannot peek into stack %d because it is empty", stackIndex))
	}

	return stack.stack[stack.stackIndexes[stackIndex]], nil
}

func (stack *MultiArrayStack) IsEmpty(stackIndex int) bool {
	stack.AssertStackIndexIsWithinRange(stackIndex)
	return stack.stackIndexes[stackIndex] == getDefaultIndexes()[stackIndex]
}

func (stack *MultiArrayStack) AssertStackIndexIsWithinRange(stackIndex int) {
	if stackIndex < 0 || stackIndex > 2 {
		panic(fmt.Sprintf("The stack index = %d is out of range", stackIndex))
	}
}
