package chapter5

func GetNextNumberWithSame1Bits(num int) (result int) {
	c := num
	c0 := 0
	c1 := 0

	for c&1 == 0 && c != 0 {
		c0++
		c >>= 1
	}

	for c&1 == 1 {
		c1++
		c >>= 1
	}

	// Error if n == 11..1100..00 then there's no bigger number with the same number of 1s
	if (c0+c1 == 31) || c0+c1 == 0 {
		return -1
	}

	p := c0 + c1 // Position of  rightmost non-trailing zero

	num |= 1 << p              // Flip rightmost non-trailing zero
	num &= ^((1 << p) - 1)     // Clear all bits to the right of p
	num |= (1 << (c1 - 1)) - 1 // Insert c-1 ones on the right

	return num
}

func GetPreviousNumberWithSame1Bits(num int) (result int) {
	temp := num
	c0 := 0
	c1 := 0

	for temp&1 == 1 {
		c1++
		temp >>= 1
	}

	if temp == 0 {
		return -1
	}

	for (temp&1) == 0 && temp != 0 {
		c0++
		temp >>= 1
	}

	p := c0 + c1

	num &= (^0) << (p + 1)

	mask := 1<<(c1+1) - 1

	num |= mask << (c0 - 1)

	return num
}
