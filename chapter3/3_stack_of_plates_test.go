package chapter3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiDimensionalStack_Push(t *testing.T) {
	var stack MultiDimensionalStack

	for i := 0; i < 10 ;i++  {
		stack.Push(i)
	}

	t.Run("Push uses stacks with limit = 3", func(t *testing.T) {
		topNode := stack.top
		count := 0;
		for topNode != nil {
			count++
			topNode = topNode.next
		}
		assert.Equal(t, 4, count)
	})
}

func TestMultiDimensionalStack_IsEmpty(t *testing.T) {
	var stack MultiDimensionalStack
	stack.Push(1)

	t.Run("IsEmpty returns false when the stack is not empty", func(t *testing.T) {
		assert.Equal(t, false, stack.IsEmpty())
	})

	stack = MultiDimensionalStack{}
	t.Run("IsEmpty returns true when the stack is empty", func(t *testing.T) {
		assert.Equal(t, true, stack.IsEmpty())
	})
}


func TestMultiDimensionalStack_Peek(t *testing.T) {
	var stack MultiDimensionalStack

	t.Run("Peek returns the last element in the stack without removing it", func(t *testing.T) {
		stack.Push(4)

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

func TestMultiDimensionalStack_Pop(t *testing.T) {
	var stack MultiDimensionalStack

	t.Run("Pop returns the last element in the stack", func(t *testing.T) {
		stack.Push(3)
		value,_ := stack.Pop()
		assert.Equal(t, 3, value)
	})


	t.Run("Pop returns an error if the stack is empty", func(t *testing.T) {
		value,err := stack.Pop()
		assert.NotNil(t, err)
		assert.Equal(t, -1, value)
	})
}

func TestMultiDimensionalStack_PopAt(t *testing.T) {
	var stack MultiDimensionalStack

	for i := 0; i < 10;i++ {
		stack.Push(i)
	}

	val, err := stack.PopAt(2)
	assert.Nil(t, err)
	assert.Equal(t, 5, val)
}