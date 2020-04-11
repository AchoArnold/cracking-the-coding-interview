package chapter10

import "math"

// SortPeakValley sorts an array into peaks and valleys
func SortPeakValley(array []int) {
	for i := 1; i < len(array); i += 2 {
		biggestIndex := maxIndex(array, i-1, i, i+1)
		if i != biggestIndex {
			SwapArray(array, i, biggestIndex)
		}
	}
}

// SwapArray swaps an array without using a 3rd integer
func SwapArray(array []int, i, j int) {
	array[j] = array[i] + array[j]
	array[i] = array[j] - array[i]
	array[j] = array[j] - array[i]
}

func maxIndex(array []int, a, b, c int) int {
	length := len(array)

	aVal := math.MinInt32
	if a >= 0 && a < length {
		aVal = array[a]
	}

	bVal := math.MinInt32
	if b >= 0 && b < length {
		bVal = array[b]
	}

	cVal := math.MinInt32
	if c >= 0 && c < length {
		cVal = array[c]
	}

	max := int(math.Max(float64(aVal), math.Max(float64(bVal), float64(cVal))))
	if aVal == max {
		return a
	}

	if bVal == max {
		return b
	}

	return c
}
