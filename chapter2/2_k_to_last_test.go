package chapter2

import (
	"testing"
)

func TestKToLast(t *testing.T) {
	input := SinglyLinkedList{
		&Node{
			data: 1,
			next: &Node{
				data: 1,
				next: &Node{
					data: 3,
					next: &Node{
						data: 4,
						next: nil,
					},
				},
			},
		},
	}

	k := 1
	expectedData := 3

	t.Run("Testing input output", func(t *testing.T) {
		if input.kToLast(k) != expectedData {
			t.Errorf(
				"Wrong Data.\nActual Output: %d\nExpected Output: %d\n",
				input.kToLast(k),
				expectedData,
			)
		}
	})

}
