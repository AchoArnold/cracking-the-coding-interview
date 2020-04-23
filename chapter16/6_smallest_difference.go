package chapter16

import (
	"math"
	"sort"
)

// FindSmallestNonNegativeDifference returns the smallest non negative difference of any 2 elements in both arrays
func FindSmallestNonNegativeDifference(array1, array2 []int) int {
	sort.Ints(array1)
	sort.Ints(array2)

	a, b := 0, 0

	diff := math.MaxInt32

	for a < len(array1) && b < len(array2) {
		if int(math.Abs(float64(array1[a]-array2[b]))) < diff {
			diff = int(math.Abs(float64(array1[a] - array2[b])))
		}

		/* Move smaller value */
		if array1[a] < array2[b] {
			a++
		} else {
			b++
		}
	}

	return diff
}
