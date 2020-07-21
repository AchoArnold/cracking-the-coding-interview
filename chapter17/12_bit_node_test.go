package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTreeToDoublyLinkedList(t *testing.T) {
	testData := []struct {
		name   string
		tree   *BitNode
		values []int
	}{
		{
			"Valid tree",
			&BitNode{
				value: 4,
				node1: &BitNode{
					value: 2,
					node1: &BitNode{
						value: 1,
						node1: &BitNode{
							value: 0,
							node1: nil,
							node2: nil,
						},
						node2: nil,
					},
					node2: &BitNode{
						value: 3,
						node1: nil,
						node2: nil,
					},
				},
				node2: &BitNode{
					value: 5,
					node1: nil,
					node2: &BitNode{
						value: 6,
						node1: nil,
						node2: nil,
					},
				},
			},
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			"empty root",
			nil,
			[]int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			list := convertToCircular(data.tree)
			values := bitNodeLinkedListToArray(list)
			assert.ElementsMatch(t, data.values, values)
		})
	}

}

func bitNodeLinkedListToArray(head *BitNode) []int {
	node := head
	var values []int
	first := 0
	if node != nil {
		first = node.value
	}

	for node != nil {
		values = append(values, node.value)
		node = node.node2
		if node.value == first {
			break
		}
	}

	return values
}
