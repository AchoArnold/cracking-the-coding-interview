package util

import (
	"fmt"
	"strconv"
	"testing"
)

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func PrintMatrixToStdOut(inputMatrix [][]int) {
	for rowIndex := 0; rowIndex < len(inputMatrix); rowIndex++{
		print("{")
		for columnIndex := 0; columnIndex < len(inputMatrix[rowIndex]); columnIndex++ {
			fmt.Printf("%d,", inputMatrix[rowIndex][columnIndex])
		}
		print("}\n")
	}
}

func HandleError(expectedMatrix [][]int, actualMatrix [][]int, t *testing.T) {
	t.Errorf(
		"The matrices below are not equal\n\nExpected matrix:\n\n%s\nActualMatrix:\n\n%s\n",
		printMatrix(expectedMatrix),
		printMatrix(actualMatrix),
	)
}

func printMatrix(inputMatrix [][]int) string {
	outputString := ""

	for rowIndex := 0; rowIndex < len(inputMatrix); rowIndex++{
		outputString += "{"
		for columnIndex := 0; columnIndex < len(inputMatrix[rowIndex]); columnIndex++ {
			outputString += fmt.Sprintf("%d,", inputMatrix[rowIndex][columnIndex])
		}
		outputString += "}\n"
	}

	return outputString
}

func printArray(array []int) string {
	result := "{"
	for _, val := range array {
		result += strconv.Itoa(val) + ", "
	}

	return result + "}"
}

func printMultiDimensionalArray(array [][]int) string {
	result := "{\n"
	for _, val := range array {
		result += "\t" + printArray(val) + ",\n"
	}

	return result + "}"
}
