package chapter16

import "math"

// GetContiguousSequenceBruteForce returns the sum of the largest sequence of increasing elements from an array by testing all elements
func GetContiguousSequenceBruteForce(input []int) (sum int) {
	sum = math.MinInt32
	for i := 0; i < len(input); i++ {
		localSum := input[i]
		if localSum > sum {
			sum = localSum
		}

		for j := i + 1; j < len(input); j++ {
			localSum += input[j]
			if localSum > sum {
				sum = localSum
			}
		}
	}

	return sum
}

// GetContiguousSequence returns the sum of largest sequence of elements in an input array
func GetContiguousSequence(input []int) (sum int) {
	maxSum := math.MinInt32
	sum = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
		if maxSum < sum {
			maxSum = sum
		} else if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
