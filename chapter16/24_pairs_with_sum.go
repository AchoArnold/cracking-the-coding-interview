package chapter16

import "sort"

// PairsWithSum returns all pairs of numbers having a specific sum
func PairsWithSum(array []int, sum int) (result [][]int) {
	sort.Ints(array)

	first := 0
	last := len(array) - 1

	for first < last {
		s := array[first] + array[last]
		if s == sum {
			result = append(result, []int{array[first], array[last]})
			first++
			last--
		} else {
			if s < sum {
				first++
			} else {
				last--
			}

		}
	}

	return result
}
