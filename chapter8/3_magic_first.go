package chapter8

import "math"

func MagicFirstBruteForce(array []int) (i int) {
	for index, val := range array {
		if index == val {
			return index
		}
	}

	return -1
}

func MagicBinarySearch(array []int) (result int) {
	return magicBinarySearch(array, 0, len(array)-1)
}

func magicBinarySearch(array []int, start, end int) (result int) {
	if end < start {
		return -1
	}

	midIndex := (start + end) / 2
	midValue := array[midIndex]
	if midValue == midIndex {
		return midIndex
	}

	// Search Left
	leftIndex := int(math.Min(float64(midIndex-1), float64(midValue)))
	left := magicBinarySearch(array, start, leftIndex)
	if left >= 0 {
		return left
	}

	// Search Right
	rightIndex := int(math.Max(float64(midIndex+1), float64(midValue)))
	right := magicBinarySearch(array, rightIndex, end)

	return right
}
