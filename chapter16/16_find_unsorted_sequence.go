package chapter16

// FindUnsortedSequence returns the start and left of a sequence of elements in an array such that if those elements
// are sorted, then the array will be sorted.
func FindUnsortedSequence(array []int) (left, right int) {
	endLeft := findEndOfLeftSubsequence(array)
	if endLeft >= len(array)-1 {
		return -1, -1
	}

	// find right subsequence
	startRight := findStartOfRightSubSequence(array)

	// get min and max
	maxIndex := endLeft
	minIndex := startRight

	for i := endLeft + 1; i < startRight; i++ {
		if array[i] < array[minIndex] {
			minIndex = i
		}

		if array[i] > array[maxIndex] {
			maxIndex = i
		}
	}

	// slide left until less than array[minIndex]
	leftIndex := shrinkLeft(array, minIndex, endLeft)

	// slide right until greater than array[maxIndex]
	rightIndex := shrinkRight(array, maxIndex, startRight)

	return leftIndex, rightIndex
}

func findEndOfLeftSubsequence(array []int) int {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			return i - 1
		}
	}
	return len(array) - 1
}

func findStartOfRightSubSequence(array []int) int {
	for i := len(array) - 2; i >= 0; i-- {
		if array[i] > array[i+1] {
			return i + 1
		}
	}

	return 0
}

func shrinkLeft(array []int, minIndex, start int) int {
	comp := array[minIndex]
	for i := start - 1; i >= 0; i-- {
		if array[i] <= comp {
			return i + 1
		}
	}

	return 0
}

func shrinkRight(array []int, maxIndex, start int) int {
	comp := array[maxIndex]
	for i := start; i < len(array); i++ {
		if array[i] >= comp {
			return i - 1
		}
	}

	return len(array) - 1
}
