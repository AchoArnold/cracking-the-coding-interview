package chapter3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAwareStack_Pop(t *testing.T) {
	top := &StackNode{4, &StackNode{5, nil}}
	stack := MinAwareStack{top, &Stack{&StackNode{4, &StackNode{4,nil}}}}

	t.Run("Test the stack returns the element at the top of the stack", func(t *testing.T) {
		element, _ := stack.Pop()
		assert.Equal(t, 4, element)
	})

	stack = MinAwareStack{nil, nil}
	t.Run("Test the stack returns an error when the stack is empty", func(t *testing.T) {
		element, err := stack.Pop()
		assert.Equal(t, element, -1)
		assert.NotEqual(t, nil, err)
		assert.Equal(t, err.Error(), "cannot pop an empty stack")
	})
}

func TestMinAwareStack_IsEmpty(t *testing.T) {
	top := &StackNode{4, &StackNode{5, nil}}
	stack := MinAwareStack{top, &Stack{&StackNode{4, &StackNode{4,nil}}}}

	t.Run("IsEmpty returns false when the stack is not empty", func(t *testing.T) {
		assert.Equal(t, false, stack.IsEmpty())
	})

	stack = MinAwareStack{nil, nil}
	t.Run("IsEmpty returns true when the stack is empty", func(t *testing.T) {
		assert.Equal(t, true, stack.IsEmpty())
	})
}

func TestMinAwareStack_Push(t *testing.T) {
	var stack MinAwareStack
	t.Run("Push adds an item to the stack", func(t *testing.T) {
		stack.Push(2)
		element, _ := stack.Pop()
		assert.Equal(t, 2, element)
	})
}

func TestMinAwareStack_Peek(t *testing.T) {
	top := &StackNode{4, nil}
	stack := MinAwareStack{top, &Stack{&StackNode{4, nil}}}

	t.Run("Peek returns the last element in the stack without removing it", func(t *testing.T) {
		element,_ := stack.Peek()
		assert.Equal(t, 4, element)

		element,_ = stack.Pop()
		assert.Equal(t, 4, element)
	})

	t.Run("Peek returns an error if the stack is empty", func(t *testing.T) {
		_,err := stack.Peek()
		assert.NotEqual(t, nil, err)
		assert.Equal(t,"cannot peak an empty stack", err.Error())
	})
}

func TestMinAwareStack_Min(t *testing.T) {
	t.Run("Min returns the minimum element in the stack when element is in the middle", func(t *testing.T) {
		var stack MinAwareStack
		stack.Push(3)
		stack.Push(1)
		stack.Push(2)

		min, err := stack.Min()

		assert.Nil(t, err)
		assert.Equal(t, 1, min)
	})

	t.Run("Min returns the minimum element in the stack when element is in the beginning", func(t *testing.T) {
		var stack MinAwareStack
		stack.Push(1)
		stack.Push(1)
		stack.Push(2)

		min, err := stack.Min()

		assert.Nil(t, err)
		assert.Equal(t, 1, min)
	})

	t.Run("Min returns the minimum element after popping the current min", func(t *testing.T) {
		var stack MinAwareStack
		stack.Push(5)
		stack.Push(1)
		stack.Push(2)
		stack.Push(0)

		_,_ = stack.Pop()

		min, err := stack.Min()

		assert.Nil(t, err)
		assert.Equal(t, 1, min)
	})

	t.Run("Min returns the minimum element in the stack when element is in the end", func(t *testing.T) {
		var stack MinAwareStack
		stack.Push(2)
		stack.Push(1)
		stack.Push(1)

		min, err := stack.Min()

		assert.Nil(t, err)
		assert.Equal(t, 1, min)
	})

}