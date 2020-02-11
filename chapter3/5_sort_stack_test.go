package chapter3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Sort(t *testing.T) {
	var stack Stack
	t.Run("Sort can sort a stack with duplicates", func(t *testing.T) {
		for _, val := range []int{3,4,2,5,1,4}  {
			stack.Push(val)
		}
		stack.Sort()

		assert.Equal(t, []int{1,2,3,4,4,5}, stack.ToArray())
	})

	t.Run("Sort can sort a sorted stack", func(t *testing.T) {
		stack.Empty()

		for _, val := range []int{1,2,3,4,4,5}  {
			stack.Push(val)
		}
		stack.Sort()

		assert.Equal(t, []int{1,2,3,4,4,5}, stack.ToArray())
	})

}
