package chapter2

import (
	"testing"
)

func TestPartitionMethod(t *testing.T) {
	inputList := SinglyLinkedList{
		&Node{
			data: 3,
			next: &Node{
				data: 5,
				next: &Node{
					data: 8,
					next: &Node{
						data: 5,
						next: &Node{
							data: 10,
							next: &Node{
								data: 2,
								next: &Node{
									data: 1,
									next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, partition := range []int{5, 8} {
		inputList.partition(partition)
		partitionFound := false

		node := inputList.head
		for node != nil {
			if node.data >= partition {
				partitionFound = true
			}

			if node.data < partition && partitionFound {
				t.Errorf("Wrong sequence of entries %s", inputList.String())
			}

			if node.data >= partition && !partitionFound {
				t.Errorf("Wrong sequence of entries %s", inputList.String())
			}

			node = node.next
		}
	}
}
