package chapter10

// SearchRotatedArray searches a rotated array for value X
func SearchRotatedArray(array []int, x int) int {
	return searchArray(array, 0, len(array)-1, x)
}

func searchArray(a []int, left, right, x int) int {
	mid := (left + right) / 2
	if x == a[mid] {
		return mid
	}

	if right < left {
		return -1
	}

	// either the left or the right hal must be normally ordered. Find out which side is normally ordered and then use
	// the normally ordered half to figure out where to find x
	if a[left] < a[mid] { // left is normally ordered
		if x >= a[left] && x < a[mid] {
			return searchArray(a, left, mid-1, x) // search left
		}
		return searchArray(a, mid+1, right, x) // search right
	} else if a[mid] < a[left] { // Right is normally ordered
		if x > a[mid] && x <= a[right] {
			return searchArray(a, mid+1, right, x) // search right
		}
		return searchArray(a, left, mid-1, x) //search left
	} else if a[left] == a[mid] { // left or right half is all repeats
		if a[mid] != a[right] { // if right is different, search it
			return searchArray(a, mid+1, right, x) // search right
		}
		// Else, we have to search both halves
		result := searchArray(a, left, mid-1, x) // search left
		if result == -1 {
			return searchArray(a, mid+1, right, x)
		}
		return result
	}

	return -1
}
