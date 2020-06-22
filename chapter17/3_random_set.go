package chapter17

import "math/rand"

// PickMRecursively pics M elements from an array randomly
func PickMRecursively(original []int, m int) []int {
	// fill subset array with first part of original array
	subset := original[0:m]

	// go through rest of original array
	for i := m; i < len(original); i++ {
		k := rand.Intn(i)
		if k < m {
			subset[k] = original[i]
		}
	}

	return subset
}
