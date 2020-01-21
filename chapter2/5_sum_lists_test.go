package chapter2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumLists(t *testing.T) {
	list1 := &SinglyLinkedList{
		&Node{
			data: 7,
			next: &Node{
				data: 1,
				next: &Node{
					data: 6,
					next: nil,
				},
			},
		},
	}

	list2 := &SinglyLinkedList{
		&Node{
			data: 5,
			next: &Node{
				data: 9,
				next: &Node{
					data: 2,
					next: nil,
				},
			},
		},
	}

	expectedResult := []int{9, 1, 2}

	t.Run("Testing Sum", func(t *testing.T) {
		var actualResult []int
		node := sumLists(list1, list2).head

		for node != nil {
			actualResult = append(actualResult, node.data)
			node = node.next
		}

		assert.Equal(t, expectedResult, actualResult, "The expected result is not equal to the actual result")
	})

}
