package main

import (
	"fmt"
	"testing"
	"reflect"
)

type testInputOutput struct {
	inputMatrix [][]int
	outputMatrix [][]int
}

func TestRotateMatrix(t *testing.T)  {
	 data := getTestData()
	 for _, testData := range data {
	 	t.Run("Testing Input Output", func(t *testing.T) {
			if !reflect.DeepEqual(rotateMatrix(testData.inputMatrix),testData.outputMatrix) {
				handleError(testData.outputMatrix, testData.inputMatrix,t)
			}
		})
	 }
}

func TestRotateMatrixConstantSpace(t *testing.T)  {
	data := getTestData()
	for _, testData := range data {
		t.Run("Testing Input Output", func(t *testing.T) {
			if !reflect.DeepEqual(rotateMatrixConstantSpace(testData.inputMatrix),testData.outputMatrix) {
				handleError(testData.outputMatrix, testData.inputMatrix,t)
			}
		})
	}
}

func handleError(expectedMatrix [][]int, actualMatrix [][]int, t *testing.T) {
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

func getTestData() []testInputOutput {
	return []testInputOutput{
		{
			[][]int{
				{1,	2,	3,	4},
				{5,	6,	7,	8},
				{9,	0,	1,	2},
				{3,	4,	5,	6},
			},
			[][]int{
				{3,9,5,1,},
				{4,0,6,2,},
				{5,1,7,3,},
				{6,2,8,4,},
			},
		},
	}
}