package chapter16

import (
	"github.com/KyleBanks/go-kit/unique"
)

// GetDivingBoardLengths gets the length that can be made up of 2 weights
func GetDivingBoardLengths(number, shorter, longer int) []int {
	vals := getDivingBoardLengths(number, 0, shorter, longer, []int{})
	return unique.Ints(vals)
}

// GetDivingBoardLoop gets the length using 1 loop.
func GetDivingBoardLoop(number, shorter, longer int) []int {
	var result []int
	for nShorter := 0; nShorter <= number; nShorter++ {
		nLonger := number - nShorter
		length := nShorter*shorter + nLonger*longer
		result = append(result, length)
	}

	return unique.Ints(result)
}

func getDivingBoardLengths(number, total, shorter, longer int, length []int) []int {
	if number == 0 {
		return append(length, total)
	}

	length = getDivingBoardLengths(number-1, total+shorter, shorter, longer, length)
	return getDivingBoardLengths(number-1, total+longer, shorter, longer, length)
}
