package chapter17

// AddWithoutPlus adds 2 numbers without using the `+` sign
func AddWithoutPlus(a int, b int) (result int) {
	if b == 0 {
		return a
	}

	sum := a ^ b          // add without carrying
	carry := (a & b) << 1 // carry but don't add

	return AddWithoutPlus(sum, carry) // recurse with sum + carry
}
