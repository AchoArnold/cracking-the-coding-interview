package chapter2

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := SinglyLinkedList{
		&Node{
			data: 1,
			next: &Node{
				data: 1,
				next: &Node{
					data: 3,
					next: nil,
				},
			},
		},
	}

	output := SinglyLinkedList{
		&Node{
			data: 1,
			next: &Node{
				data: 3,
				next: nil,
			},
		},
	}

	t.Run("Testing input output", func(t *testing.T) {
		if input.removeDuplicates().print() != output.print() {
			t.Errorf(
				"The 2 lists below are not equal.\nActual Output: %s\nExpected Output: %s\n",
				input.print(),
				output.print(),
			)
		}
	})

}
