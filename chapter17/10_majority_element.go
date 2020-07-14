package chapter17

// FindMajorityElement finds an element in a list of elements which repeats more than 50% of the time
func FindMajorityElement(values []int) int {
	if len(values) == 0 {
		return -1
	}

	currVal := values[0]
	currCount := 1

	for i := 1; i < len(values); i++ {
		if currVal == values[i] {
			currCount++
		} else {
			currCount--

			if currCount == 0 {
				currVal = values[i]
				currCount = 1
			}
		}
	}

	currCount = 0
	for _, val := range values {
		if val == currVal {
			currCount++
		}
	}

	if currCount > len(values)/2 {
		return currVal
	}

	return -1
}
