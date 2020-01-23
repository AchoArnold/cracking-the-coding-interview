package chapter2

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestIntersection(t *testing.T) {
	node1 := &Node{nil, 1}
	node5 := &Node{nil, 5}
	node4 := &Node{node5, 4}
	node3 := &Node{node4, 3}
	node2 := &Node{node3, 2}

	testData := []struct {
		list1, list2     *SinglyLinkedList
		intersectingNode *Node
	}{
		{
			&SinglyLinkedList{node3},
			&SinglyLinkedList{node2},
			node3,
		},
		{
			&SinglyLinkedList{node1},
			&SinglyLinkedList{node2},
			nil,
		},
	}

	for index, data := range testData {
		t.Run("Dataset "+strconv.Itoa(index), func(t *testing.T) {
			assert.Equal(t, data.intersectingNode, intersection(data.list1, data.list2))
		})
	}
}
