package chapter17

// SubSquare represents a subsquare from a matrix
type SubSquare struct {
	row        int
	col        int
	squareSize int
}

// FindSquare returns the biggest square in the matrix
func FindSquare(matrix [][]int) *SubSquare {
	for i := len(matrix); i >= 1; i-- {
		square := findSquareWithSize(matrix, i)
		if square != nil {
			return square
		}
	}

	return nil
}

func findSquareWithSize(matrix [][]int, squareSize int) *SubSquare {
	// ON an edge of length N, there are (N - sz + 1) squares of length sz
	count := len(matrix) - squareSize + 1

	// Iterate through all squares with side length squareSize
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if isSquare(matrix, row, col, squareSize) {
				return &SubSquare{row, col, squareSize}
			}
		}
	}

	return nil
}

func isSquare(matrix [][]int, row, col, size int) bool {
	// Check top and bottom border.
	for j := 0; j < size; j++ {
		if matrix[row][col+j] == 1 {
			return false
		}

		if matrix[row+size-1][col+j] == 1 {
			return false
		}
	}

	// check left and right border
	for i := 0; i < size-1; i++ {
		if matrix[row+i][col] == 1 {
			return false
		}

		if matrix[row+i][col+size-1] == 1 {
			return false
		}
	}

	return true
}
