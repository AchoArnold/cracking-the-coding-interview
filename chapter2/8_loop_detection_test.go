package chapter2

import (
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoopDetection(t *testing.T) {
	node1 := &Node{nil, 1}
	node5 := &Node{nil, 5}
	node4 := &Node{node5, 4}
	node3 := &Node{node4, 3}
	node2 := &Node{node3, 2}
	node5.next = node3

	testData := []struct {
		list        *SinglyLinkedList
		loopingNode *Node
	}{
		{
			&SinglyLinkedList{node2},
			node3,
		},
		{
			&SinglyLinkedList{node1},
			nil,
		},
	}

	for index, data := range testData  {
		t.Run("Test Data " + strconv.Itoa(index), func(t *testing.T) {
			assert.Equal(t, data.loopingNode, data.list.GetLoopingNode())
		})
	}
}
