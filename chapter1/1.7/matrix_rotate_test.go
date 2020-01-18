package main

import (
	"testing"
	"reflect"
	util "github.com/AchoArnold/cracking-the-coding-interview/utils"
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
				util.handleError(testData.outputMatrix, testData.inputMatrix,t)
			}
		})
	 }
}

func TestRotateMatrixConstantSpace(t *testing.T)  {
	data := getTestData()
	for _, testData := range data {
		t.Run("Testing Input Output", func(t *testing.T) {
			if !reflect.DeepEqual(rotateMatrixConstantSpace(testData.inputMatrix),testData.outputMatrix) {
				util.handleError(testData.outputMatrix, testData.inputMatrix,t)
			}
		})
	}
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