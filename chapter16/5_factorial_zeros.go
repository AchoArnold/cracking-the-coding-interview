package chapter16

// CountFactorialZeros counts the number of zeros in the factorial of num
func CountFactorialZeros(num int) int {
	if num <= 0 {
		return -1
	}

	count := 0
	for i := 5; num/i > 0; i *= 5 {
		count += num / i
	}
	return count
}
