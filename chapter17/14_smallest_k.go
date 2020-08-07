package chapter17

import "math/rand"

// partitionResult stores the result of the partition
type partitionResult struct {
	leftSize   int
	middleSize int
}

func newPartitionResult(left, middle int) partitionResult {
	return partitionResult{
		leftSize:   left,
		middleSize: middle,
	}
}

// SmallestK returns the smallest K items in an array
func SmallestK(array []int, k int) []int {
	if k <= 0 || k > len(array) {
		return []int{}
	}

	// get item with rank k-1
	threshold := rank(array, k-1)

	// copy elements smaller than the  threshold element
	smallest := make([]int, k)
	count := 0
	for _, a := range array {
		if a < threshold {
			smallest[count] = a
			count++
		}
	}

	// if there's still room left, this must be for elements equal to the threshold element. copy those in
	for count < k {
		smallest[count] = threshold
		count++
	}

	return smallest
}

// find value with rank k in array
func rank(array []int, k int) int {
	return rankStartEnd(array, k, 0, len(array)-1)
}

// find value with rank k in sub array between start and end
func rankStartEnd(array []int, k, start, end int) int {
	// partition array around an arbitrary point
	pivot := array[rand.Intn(end)]
	partition := partition(array, start, end, pivot)
	leftSize := partition.leftSize
	middleSize := partition.middleSize

	// Search portion of array
	if k < leftSize {
		// rank k is on the left half
		return rankStartEnd(array, k, start, start+leftSize-1)
	}

	if k < leftSize+middleSize {
		// rank k i s in middle
		// middle is all pivot values
		return pivot
	}

	// rank k is on the right
	return rankStartEnd(array, k-leftSize-middleSize, start+leftSize+middleSize, end)
}

func arraySwap(array []int, i, j int) {
	t := array[i]
	array[i] = array[j]
	array[j] = t
}

// Partition result into < pivot, equal to pivot -> bigger than pivot
func partition(array []int, start, end, pivot int) partitionResult {
	left := start   // stays at (right) edge of left side
	right := end    // stays at (left) edge of right side
	middle := start // stays at (right) edge of middle

	for middle < right {
		if array[middle] < pivot {
			// middle is smaller than the pivot. Left is either smaller or equal the pivot
			// Either way, swap them. Then middle and left should move by one
			arraySwap(array, middle, left)
			middle++
			left++
		} else if array[middle] > pivot {
			// Middle is bigger than the pivot. Right could have any value.
			// Swap them, then we know that the new right is bigger than the pivot.
			// Move right by one
			arraySwap(array, middle, right)
			right--
		} else if array[middle] == pivot {
			// middle is equal to the pivot. Move by one
			middle++
		}
	}

	// return sizes of left and middle
	return newPartitionResult(left-start, right-left+1)
}
