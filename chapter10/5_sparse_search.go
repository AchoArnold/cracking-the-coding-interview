package chapter10

import "math"

// SearchStringExcludingEmpty searches for a string excluding the empty string
func SearchStringExcludingEmpty(strings []string, str string) (index int) {
	if len(strings) == 0 {
		return -1
	}

	return searchStringsExcludingEMpty(strings, str, 0, len(strings)-1)
}

func searchStringsExcludingEMpty(strings []string, str string, first, last int) int {
	if first > last {
		return -1
	}

	mid := int(math.Ceil(float64(first+last) / float64(2)))

	if len(strings[mid]) == 0 {
		left := mid - 1
		right := mid + 1

		for {
			if left < first && right > last {
				return -1
			}

			if right <= last && len(strings[right]) != 0 {
				mid = right
				break
			}

			if left >= first && len(strings[first]) != 0 {
				mid = left
				break
			}

			right++
			left--
		}
	}

	// check for string and recurse if necessary
	if strings[mid] == str {
		return mid
	}

	if strings[mid] < str {
		return searchStringsExcludingEMpty(strings, str, mid+1, last) // search right
	}

	return searchStringsExcludingEMpty(strings, str, first, mid-1) // search left
}
