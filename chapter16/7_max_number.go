package chapter16

// MaxNumWithoutComparison gets the max of 2 numbers without comparing
func MaxNumWithoutComparison(a, b int) int {
	sa, sb, sc := sign(a), sign(b), sign(a-b)

	// goal: define a value of k which is 1 if a  > b and 0 if a < b
	// if a == b it doesn't matter what value k is

	// if a and b have different signs then k = sign(a)
	useSignOfA := sa ^ sb

	// if a and b have the same sign then k = sign(a - b)
	useSignOfC := flip(sa ^ sb)

	k := useSignOfA*sa + useSignOfC*sc
	q := flip(k) //opposite of k

	return a*k + b*q
}

// flip 0 to 1 and 1 to zero
func flip(bit int) int {
	return 1 ^ bit
}

func sign(a int) int {
	return flip((a >> 31) & 0x1)
}
