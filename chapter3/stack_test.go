package chapter3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Pop(t *testing.T) {
	stack := Stack{&StackNode{4, &StackNode{5, nil}}}
	t.Run("Test the stack returns the element at the top of the stack", func(t *testing.T) {
		element, _ := stack.Pop()
		assert.Equal(t, 4, element)
	})

	stack = Stack{nil}
	t.Run("Test the stack returns an error when the stack is empty", func(t *testing.T) {
		element, err := stack.Pop()
		assert.Equal(t, element, -1)
		assert.Equal(t, err.Error(), "cannot pop an empty stack")
	})
}

func TestStack_IsEmpty(t *testing.T) {
	stack := Stack{&StackNode{1, nil}}
	t.Run("IsEmpty returns false when the stack is not empty", func(t *testing.T) {
		assert.Equal(t, false, stack.IsEmpty())
	})

	stack = Stack{nil}
	t.Run("IsEmpty returns true when the stack is empty", func(t *testing.T) {
		assert.Equal(t, true, stack.IsEmpty())
	})
}

func TestStack_Push(t *testing.T) {
	var stack Stack
	t.Run("Push adds an item to the stack", func(t *testing.T) {
		stack.Push(2)
		element, _ := stack.Pop()
		assert.Equal(t, 2, element)
	})
}

func TestStack_Peek(t *testing.T) {
	stack := Stack{&StackNode{3, nil}}

	t.Run("Peek returns the last element in the stack without removing it", func(t *testing.T) {
		element, _ := stack.Peek()
		assert.Equal(t, 3, element)

		element, _ = stack.Pop()
		assert.Equal(t, 3, element)
	})

	t.Run("Peek returns an error if the stack is empty", func(t *testing.T) {
		_, err := stack.Peek()
		assert.Equal(t, "cannot peak an empty stack", err.Error())
	})
}
