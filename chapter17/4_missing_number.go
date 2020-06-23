package chapter17

// FindMissing returns the missing item from a list of integers from [0,n]
func FindMissing(array []int) int {
	return findMissing(array, 0)
}

func findMissing(input []int, column int) int {
	if column >= 32 {
		return 0
	}

	var oneBits []int
	var zeroBits []int

	for _, t := range input {
		if fetchBit(t, column) == 0 {
			zeroBits = append(zeroBits, t)
		} else {
			oneBits = append(oneBits, t)
		}
	}

	if len(zeroBits) <= len(oneBits) {
		v := findMissing(zeroBits, column+1)
		return (v << 1) | 0
	}

	v := findMissing(oneBits, column+1)
	return (v << 1) | 1

}

func fetchBit(num, col int) int {
	return (num >> col) & 1
}
