package chapter5

func SetBitsInRange(n, m,i,j int) (result int) {
	// Create a mask to clear bits i through j in n
	allOnes := ^0

	// 1s before position j, then 0s.
	left := allOnes << (j + 1)

	// 1's after position i
	right := (1 <<  i) -1

	// All 1s except for 0s between i and j
	mask := left | right

	// Clear bits j through i then put m in there
	nCleared := n & mask
	mShifted := m << i

	// perform an or
	return nCleared | mShifted
}
