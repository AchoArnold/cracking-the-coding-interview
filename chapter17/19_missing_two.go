package chapter17

import (
	"math"
)

// MissingTwo returns the 2 missing numbers of the array ordered from 0 to N
func MissingTwo(array []int) []int {
	maxValue := len(array) + 2
	remSquare := squareSumToN(maxValue)
	remOne := maxValue * (maxValue + 1) / 2

	for i := 0; i < len(array); i++ {
		remSquare -= array[i] * array[i]
		remOne -= array[i]
	}

	return solveEquation(remOne, remSquare)
}

func solveEquation(r1, r2 int) []int {
	// ax^2 + bx^2 + c
	// x = [-b +- sqrt(b^2 - 4ac)]/2a
	// In this case it has to be a + and not a -
	a := 2
	b := -2 * r1
	c := r1*r1 - r2

	part1 := float64(-1 * b)
	part2 := math.Sqrt(float64(b*b - 4*a*c))
	part3 := float64(2 * a)

	solutionX := int((part1 + part2) / part3)
	solutionY := r1 - solutionX

	return []int{solutionX, solutionY}
}

func squareSumToN(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}
