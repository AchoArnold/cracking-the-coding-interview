package chapter2

func sumLists(list1 *SinglyLinkedList, list2 *SinglyLinkedList) *SinglyLinkedList {
	var resultList SinglyLinkedList
	list1CurrentNode := list1.head
	list2CurrentNode := list2.head
	carry := 0

	for list1CurrentNode != nil || list2CurrentNode != nil {
		list1Value, list2Value := 0, 0

		if list1CurrentNode != nil {
			list1Value = list1CurrentNode.data
		}

		if list2CurrentNode != nil {
			list2Value = list2CurrentNode.data
		}

		sumResult := carry + list1Value + list2Value
		if sumResult >= 10 {
			sumResult = sumResult % 10
			carry = 1
		}

		if (resultList == SinglyLinkedList{}) {
			resultList = SinglyLinkedList{
				&Node{
					next: nil,
					data: sumResult,
				},
			}
		} else {
			resultList.appendToHead(sumResult)
		}

		if list1CurrentNode != nil {
			list1CurrentNode = list1CurrentNode.next
		}

		if list2CurrentNode != nil {
			list2CurrentNode = list2CurrentNode.next
		}
	}

	return &resultList
}

func sumListsReverse(list1 *SinglyLinkedList, list2 *SinglyLinkedList) *SinglyLinkedList {
	return sumLists(list1.Reverse(), list2.Reverse())
}
