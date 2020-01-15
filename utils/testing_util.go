package util

import "fmt"

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func PrintMatrix(inputMatrix [][]int) {
	for rowIndex := 0; rowIndex < len(inputMatrix); rowIndex++{
		print("{")
		for columnIndex := 0; columnIndex < len(inputMatrix[rowIndex]); columnIndex++ {
			fmt.Printf("%d,", inputMatrix[rowIndex][columnIndex])
		}
		print("}\n")
	}
}