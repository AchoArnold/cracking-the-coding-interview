package chapter10

import "math"

// MergeSort sorts an item recursively using the mergesort algorithm
func MergeSort(items []int) {
	cache := make([]int, len(items))
	mergeSortRecursive(items, cache, 0, len(items)-1)
}

func mergeSortRecursive(items []int, cache []int, start, end int) {
	if start < end {
		middle := int(math.Floor(float64(start+end) / float64(2)))
		mergeSortRecursive(items, cache, start, middle)
		mergeSortRecursive(items, cache, middle+1, end)
		merge(items, cache, start, middle, end)
	}
}

func merge(items []int, cache []int, start, middle, end int) {
	for i := start; i <= end; i++ {
		cache[i] = items[i]
	}

	helperLeft := start
	helperRight := middle + 1
	current := start

	for helperLeft <= middle && helperRight <= end {
		if cache[helperLeft] <= cache[helperRight] {
			items[current] = cache[helperLeft]
			helperLeft++
		} else {
			items[current] = cache[helperRight]
			helperRight++
		}
		current++
	}

	// Copy the rest of the left side of the array into the targer array
	remaining := middle - helperLeft
	for i := 0; i <= remaining; i++ {
		items[current+i] = cache[helperLeft+i]
	}
}
