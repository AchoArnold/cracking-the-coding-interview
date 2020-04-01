package main

import (
	"reflect"
	"testing"

	util "github.com/AchoArnold/cracking-the-coding-interview/utils"
)

func TestZeroMatrix(t *testing.T) {
	inputOutput := []struct {
		input          [][]int
		expectedOutput [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			[][]int{
				{1, 2, 0},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{7, 0, 0},
			},
		},
	}

	for _, testData := range inputOutput {
		t.Run("Testing", func(t *testing.T) {
			if reflect.DeepEqual(zeroMatrix(testData.input), testData.expectedOutput) == false {
				util.HandleError(testData.expectedOutput, testData.input, t)
			}
		})
	}
}
