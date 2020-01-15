package main

func rotateMatrix(matrix [][]int) [][]int {
	matrixSize := len(matrix)

	newMatrix := make([][]int, matrixSize)
	for index := 0; index < matrixSize ; index++  {
		newMatrix[index] = make([]int, matrixSize)
	}

	for rowIndex := 0; rowIndex < matrixSize; rowIndex++{
		for columnIndex := 0; columnIndex < matrixSize; columnIndex++ {
			matrixEntry := matrix[rowIndex][columnIndex]
			newRowIndex := matrixSize - rowIndex -1
			newMatrix[columnIndex][newRowIndex] = matrixEntry
		}

	}

	return newMatrix
}


// Transpose then row reverse
func rotateMatrixConstantSpace(matrix [][]int) [][]int {
	return reverseRows(transposeMatrix(matrix))
}

func transposeMatrix(matrix [][]int) [][]int {
	matrixSize := len(matrix)
	columnStart := 0

	for rowIndex := 0; rowIndex < matrixSize; rowIndex++{
		for columnIndex := columnStart; columnIndex < matrixSize; columnIndex++ {
			matrixEntry := matrix[rowIndex][columnIndex]
			transposeEntry := matrix[columnIndex][rowIndex]

			matrix[rowIndex][columnIndex] = transposeEntry
			matrix[columnIndex][rowIndex] = matrixEntry
		}

		columnStart++
	}

	return matrix
}

func reverseRows(matrix [][]int) [][]int {
	matrixSize := len(matrix)
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++{
		for columnIndex := 0; columnIndex < matrixSize/2; columnIndex++ {
			matrixEntry := matrix[rowIndex][columnIndex]

			newColumnIndex := matrixSize - columnIndex - 1
			flippedEntry := matrix[rowIndex][newColumnIndex]

			matrix[rowIndex][columnIndex] = flippedEntry
			matrix[rowIndex][newColumnIndex] = matrixEntry
		}
	}

	return matrix
}