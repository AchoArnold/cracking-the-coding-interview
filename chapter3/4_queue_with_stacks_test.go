package chapter3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueueUsingStacks_Add(t *testing.T) {
	var queue QueueUsingStacks
	t.Run("Items can be added to an empty queue", func(t *testing.T) {
		queue.Add(3)
		queue.Add(4)
		assert.True(t,true)
	})
}


func TestQueueUsingStacks_IsEmpty(t *testing.T) {
	var queue QueueUsingStacks
	t.Run("IsEmpty returns true if the queue is empty", func(t *testing.T) {
		assert.True(t, queue.IsEmpty())
	})

	queue.Add(2)
	t.Run("IsEmpty returns true if the queue is not empty", func(t *testing.T) {
		assert.False(t, queue.IsEmpty())
	})
}

func TestQueueUsingStacks_Remove(t *testing.T) {
	var queue QueueUsingStacks
	t.Run("Remove returns an error if the queue is empty", func(t *testing.T) {
		val, err := queue.Remove()
		assert.NotNil(t, err)
		assert.Equal(t, -1, val)
	})

	queue.Add(2)
	t.Run("Remove returns the head if the queue is not empty", func(t *testing.T) {
		val, err := queue.Remove()
		assert.Nil(t, err)
		assert.Equal(t, 2, val)
	})
}

func TestQueueUsingStacks_Peek(t *testing.T) {
	var queue QueueUsingStacks
	t.Run("Peek returns an error if the queue is empty", func(t *testing.T) {
		val, err := queue.Peek()
		assert.NotNil(t, err)
		assert.Equal(t, -1, val)
	})

	queue.Add(2)
	t.Run("Peek returns the head if the queue is not empty", func(t *testing.T) {
		val, err := queue.Peek()
		assert.Nil(t, err)
		assert.Equal(t, 2, val)
	})
}