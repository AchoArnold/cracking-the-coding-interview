package chapter2

func (linkedList *SinglyLinkedList) IsPalindrome() bool {
	newList := linkedList.Clone()

	listAsSlice := newList.Slice()
	listReverseSlice := newList.Reverse().Slice()

	for index := 0; index < len(listAsSlice); index++ {
		if listAsSlice[index] != listReverseSlice[index] {
			return false
		}
	}

	return true
}
