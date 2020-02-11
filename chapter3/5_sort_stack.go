package chapter3

func (stack *Stack) Sort() {
	if stack.IsEmpty() || stack.top.next == nil {
		return
	}

	for {
		secNode := stack.top
		valChanged := false

		for secNode.next != nil {
			if secNode.data > secNode.next.data {
				val := secNode.data
				secNode.data = secNode.next.data
				secNode.next.data = val
				valChanged = true
			}

			secNode = secNode.next
		}

		if valChanged == false {
			break
		}
	}
}
