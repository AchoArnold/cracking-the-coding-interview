package chapter17

import "unicode"

// FindLongestSubarray returns the longest subarray with equal letters and numbers
func FindLongestSubarray(array []rune) []rune {
	// compute deltas between count of numbers and count of letters
	deltas := computeDeltaArray(array)

	// find pair in deltas with matching values and largest span
	match := findLongestMatch(deltas)

	return array[match[0]+1 : match[1]+1]
}

func findLongestMatch(deltas []int) (match [2]int) {
	cache := map[int]int{0: -1}

	var max [2]int

	for i := 0; i < len(deltas); i++ {
		if _, ok := cache[deltas[i]]; !ok {
			cache[deltas[i]] = i
		} else {
			match := cache[deltas[i]]
			distance := i - match
			longest := max[1] - max[0]
			if distance > longest {
				max[1] = i
				max[0] = match
			}
		}
	}

	return max
}
func computeDeltaArray(array []rune) (deltas []int) {
	delta := 0
	for i := 0; i < len(array); i++ {
		if unicode.IsDigit(array[i]) {
			delta++
		} else if unicode.IsLetter(array[i]) {
			delta--
		}
		deltas = append(deltas, delta)
	}

	return deltas
}
