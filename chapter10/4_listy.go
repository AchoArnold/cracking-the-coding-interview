package chapter10

import "math"

// Listy is an array like datatype which returns -1 if an element doesn't exist
type Listy map[int]int

func (list Listy) elementAt(index int) int {
	if val, ok := list[index]; ok {
		return val
	}
	return -1
}

// SearchListy searches a a listy for a value
func SearchListy(list Listy, value int) int {
	index := 1
	for list.elementAt(index) != -1 && list.elementAt(index) < value {
		index *= 2
	}

	return listyBinarySearch(list, value, index/2, index)
}

func listyBinarySearch(list Listy, value, low, high int) int {
	for low <= high {
		mid := int(math.Ceil(float64(low+high) / float64(2)))
		middle := list.elementAt(mid)

		if middle > value || middle == -1 {
			high = mid - 1
		} else if middle < value {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
