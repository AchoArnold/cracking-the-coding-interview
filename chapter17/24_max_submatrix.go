package chapter17

// SubMatrix represents a sub matrix
type SubMatrix struct {
	row1 int
	col1 int
	row2 int
	col2 int
	sum  int
}

// MaxMatrix returns the max sub matrix
func MaxMatrix(matrix [][]int) *SubMatrix {
	rowCount := len(matrix)
	colCount := len(matrix[0])
	var bestMatrix *SubMatrix

	for row1 := 0; row1 < rowCount; row1++ {
		for row2 := row1; row2 < rowCount; row2++ {
			for col1 := 0; col1 < colCount; col1++ {
				for col2 := col1; col2 < colCount; col2++ {
					sum := sumMatrix(matrix, row1, col1, row2, col2)
					if bestMatrix == nil || bestMatrix.sum < sum {
						bestMatrix = &SubMatrix{row1, col1, row2, col2, sum}
					}
				}
			}
		}
	}

	return bestMatrix
}

func sumMatrix(matrix [][]int, row1, col1, row2, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			sum += matrix[r][c]
		}
	}

	return sum
}
