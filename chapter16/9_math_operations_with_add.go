package chapter16

import "math"

// MultiplyWithOnlyAddition multiplies two numbers using only addition
func MultiplyWithOnlyAddition(x, y int) int {
	isNegative := false
	if x < 0 {
		x = FlipSignWithAddition(x)
		isNegative = !isNegative
	}

	if y < 0 {
		y = FlipSignWithAddition(y)
		isNegative = !isNegative
	}

	min, max := x, y
	if y < x {
		min, max = y, x
	}

	result := 0
	for i := 0; i < min; i++ {
		result += max
	}

	if isNegative {
		result = FlipSignWithAddition(result)
	}

	return result
}

// SubtractWithOnlyAddition implements x - y
func SubtractWithOnlyAddition(x, y int) int {
	return x + FlipSignWithAddition(y)
}

// DivideWithOnlyAddition implements x/y
func DivideWithOnlyAddition(x, y int) int {
	if y == 0 {
		return math.MaxInt32
	}

	isNegative := false
	if x < 0 {
		x = FlipSignWithAddition(x)
		isNegative = !isNegative
	}

	if y < 0 {
		y = FlipSignWithAddition(y)
		isNegative = !isNegative
	}

	if y > x {
		return 0
	}

	count := 0
	result := y
	for result <= x {
		count++
		result = result * y
	}

	if isNegative {
		return FlipSignWithAddition(count)
	}

	return count
}

// FlipSignWithAddition flips the sign of a number using only addition.
func FlipSignWithAddition(num int) int {
	if num < 0 {
		result := 0
		count := 0
		for result > num {
			result += -1
			count++
		}

		return count
	}

	result := 0
	count := 0
	for count < num {
		result += -1
		count++
	}

	return result
}
