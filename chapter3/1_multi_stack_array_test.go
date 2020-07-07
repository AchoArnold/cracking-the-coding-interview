package chapter3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiArrayStack_Pop(t *testing.T) {
	stack := MultiArrayStack{
		[3]int{1, 5, 9},
		[12]int{4, 0, 0, 0, 3, 0, 0, 0, 2, 0, 0, 0},
	}

	for index, val := range []int{4, 3, 2} {
		t.Run("Test the stack returns the element at the top of the stack", func(t *testing.T) {
			element, _ := stack.Pop(index)
			assert.Equal(t, val, element)
		})
	}

	for index := range []int{4, 3, 2} {
		t.Run("Test the stack returns an error when the stack is empty", func(t *testing.T) {
			element, err := stack.Pop(index)
			assert.Equal(t, element, -1)
			assert.Equal(t, err.Error(), fmt.Sprintf("Cannot pop from stack %d because it is empty", index))
		})
	}

	for _, val := range []int{3, 4, -1} {
		t.Run("Test a panic is encountered if the stack index is out of range", func(t *testing.T) {
			assert.Panics(t, func() { stack.Pop(val) })
		})
	}

}

func TestMultiArrayStack_IsEmpty(t *testing.T) {
	stack := MultiArrayStack{
		[3]int{1, 5, 9},
		[12]int{4, 0, 0, 0, 3, 0, 0, 0, 2, 0, 0, 0},
	}
	for index := range []int{4, 3, 2} {
		t.Run("IsEmpty returns false when the stack is not empty", func(t *testing.T) {
			assert.Equal(t, false, stack.IsEmpty(index))
		})
	}

	stack = *NewMultiArrayStack()
	for index := range []int{4, 3, 2} {
		t.Run("IsEmpty returns true when the stack is empty", func(t *testing.T) {
			assert.Equal(t, true, stack.IsEmpty(index))
		})
	}

	for _, val := range []int{3, 4, -1} {
		t.Run("IsEmpty a panics when the stack index is out of range", func(t *testing.T) {
			assert.Panics(t, func() { stack.IsEmpty(val) })
		})
	}
}

func TestMultiArrayStack_Push(t *testing.T) {
	var stack MultiArrayStack

	for index, val := range []int{4, 3, 2} {
		t.Run("Push adds an item to the stack", func(t *testing.T) {
			err := stack.Push(index, val)
			assert.Equal(t, nil, err)

			element, err := stack.Pop(index)
			assert.Equal(t, val, element)
			assert.Equal(t, nil, err)
		})
	}

	stack = MultiArrayStack{
		[3]int{4, 8, 12},
		[12]int{4, 0, 0, 0, 3, 0, 0, 0, 2, 0, 0, 0},
	}
	for index, val := range []int{4, 3, 2} {
		t.Run("Push throws an error if the stack is full", func(t *testing.T) {
			err := stack.Push(index, val)
			assert.NotEqual(t, nil, err)
			assert.Equal(t, fmt.Sprintf("Cannot push into stack %d because it is full", index), err.Error())
		})
	}

	for _, val := range []int{3, 4, -1} {
		t.Run("Push panics when the stack index is out of range", func(t *testing.T) {
			assert.Panics(t, func() { stack.Push(val, val) })
		})
	}
}

func TestMultiArrayStack_AssertStackIndexIsWithinRange(t *testing.T) {
	var stack MultiArrayStack
	for _, val := range []int{3, 4, -1} {
		t.Run("AssertStackIndexIsWithinRange a panics when the stack index is out of range", func(t *testing.T) {
			assert.Panics(t, func() { stack.AssertStackIndexIsWithinRange(val) })
		})
	}
}
