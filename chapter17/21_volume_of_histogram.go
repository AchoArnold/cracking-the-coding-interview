package chapter17

// ComputeHistogramVolume goes through each bar and compute the volume of water above it. volume of water at the bar =
//   	height - min(tallest bar on left, tallest bar on right)
// 		[where above equation is positive]
// Compute the left max in the first sweep, then sweep again to compute the right max, minimum of the bar heights, and the delta
func ComputeHistogramVolume(histogram []int) int {
	if len(histogram) == 0 {
		return 0
	}

	// Get left max
	leftMaxes := make([]int, len(histogram))
	leftMax := histogram[0]

	for i := 0; i < len(histogram); i++ {
		leftMax = max(leftMax, histogram[i])
		leftMaxes[i] = leftMax
	}

	sum := 0

	// get right max
	rightMax := histogram[len(histogram)-1]
	for i := len(histogram) - 1; i >= 0; i-- {
		rightMax = max(rightMax, histogram[i])
		secondTallest := min(rightMax, leftMaxes[i])

		// If there are taller things on the left and right side, then there is water above this bar
		// Compute the volume and add to the sum
		if secondTallest > histogram[i] {
			sum += secondTallest - histogram[i]
		}
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
