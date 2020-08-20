package chapter17

import "math"

// MaxMinutesRecursive returns the max sequence with skipping.
func MaxMinutesRecursive(messages []int) int {
	return maxMinutesRecursive(messages, 0)
}

func maxMinutesRecursive(messages []int, index int) int {
	if index >= len(messages) {
		return 0
	}

	// best with this reservation
	bestWith := messages[index] + maxMinutesRecursive(messages, index+2)

	// best without this reservation
	bestWithout := maxMinutesRecursive(messages, index+1)

	// return best of this subarray, starting fro index
	if bestWith > bestWithout {
		return bestWith
	}

	return bestWithout
}

// MaxMinutesIterative returns the max sequence without skipping
func MaxMinutesIterative(messages []int) int {
	oneAway := 0
	twoAway := 0

	for i := len(messages) - 1; i >= 0; i-- {
		bestWith := messages[i] + twoAway
		bestWithout := oneAway

		current := int(math.Max(float64(bestWith), float64(bestWithout)))
		twoAway = oneAway
		oneAway = current
	}

	return oneAway
}
