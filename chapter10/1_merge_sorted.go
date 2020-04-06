package chapter10

// MergeSorted merges 2 sorted arrays into the first array.
func MergeSorted(a, b []int, lastA, lastB int) {
	indexA := lastA - 1 // Index of the last element in array a
	indexB := lastB - 1 // Index of the last element in array b
	indexMerged := lastB + lastA - 1

	// Merge a and b starting from the last element in each
	for indexB >= 0 {
		// end of a is > than end of b
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA] // copy element
			indexA--
		} else {
			a[indexMerged] = b[indexB] // copy element
			indexB--
		}

		indexMerged--
	}
}
