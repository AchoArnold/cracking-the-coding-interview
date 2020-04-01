package main

func zeroMatrix(matrix [][]int) [][]int {
	rowSize := len(matrix)
	colSize := len(matrix[0])

	type MatrixElement = struct {
		rowIndex int
		colIndex int
	}

	var zerosFound []MatrixElement

	for rowIndex := 0; rowIndex < rowSize; rowIndex++ {
		for colIndex := 0; colIndex < colSize; colIndex++ {
			if matrix[rowIndex][colIndex] == 0 {
				zerosFound = append(zerosFound, MatrixElement{rowIndex, colIndex})
				break
			}

		}
	}

	for _, matrixEntry := range zerosFound {

		for startIndex := 0; startIndex < rowSize; startIndex++ {
			matrix[startIndex][matrixEntry.colIndex] = 0
		}

		for startIndex := 0; startIndex < colSize; startIndex++ {
			matrix[matrixEntry.rowIndex][startIndex] = 0
		}
	}

	return matrix
}
