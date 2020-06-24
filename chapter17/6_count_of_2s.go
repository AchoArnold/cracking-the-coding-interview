package chapter17

import (
	"math"
	"strconv"
)

// Count2sInRange counts the number of 2s in a digit
func Count2sInRange(number int) int {
	count := 0
	length := len(strconv.Itoa(number))

	for digit := 0; digit < length; digit++ {
		count += count2sInRangeAtDigit(number, digit)
	}

	return count
}

func count2sInRangeAtDigit(number int, d int) int {
	powerOf10 := int(math.Pow(float64(10), float64(d)))
	nextPowerOf10 := powerOf10 * 10
	right := number % powerOf10

	roundDown := number - number%nextPowerOf10
	roundUp := roundDown + nextPowerOf10

	digit := (number / powerOf10) % 10
	if digit < 2 { //if the digit in spot digit is
		return roundDown / 10
	} else if digit == 2 {
		return roundDown/10 + right + 1
	}

	return roundUp / 10
}
