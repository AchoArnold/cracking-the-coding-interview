package chapter8

import "math"

// PlaceNQueens returns the number of ways of placing n queens in an n x n grid
func PlaceNQueens(queens int) int {
	return len(placeQueens(queens, 0, make([]int, queens), [][]int{}))
}

func placeQueens(gridSize int, row int, columns []int, results [][]int) [][]int {
	// A valid placement is found
	if row == gridSize {
		results = append(results, columns)
	} else {
		for col := 0; col < gridSize; col++ {
			if checkValid(columns, row, col) {
				columns[row] = col // place queen
				results = placeQueens(gridSize, row+1, columns, results)
			}
		}
	}

	return results
}

// checkValid checks if (row1, col1) is a valid spot for a queen by checking if there is a
// queen in the same column or diagonal. We don't nee ot check if for queens in the same row
// because the calling PlaceQueen only attempts to place one queen at a time. We know this
// row is emtpy
func checkValid(columns []int, row1, column1 int) bool {
	for row2 := 0; row2 < row1; row2++ {
		column2 := columns[row2]

		// Check if (row2, column2) invalidates (row1, column1 as a queen spot)

		// Check if rows have a queen in the same column
		if column1 == column2 {
			return false
		}

		// check diagonals: If the distance between the columns equals the distance between row,
		// then the're in the same diagonal
		columnDistance := int(math.Abs(float64(column2 - column1)))

		// row1 > row2 so no need for abs
		rowDistance := row1 - row2

		if columnDistance == rowDistance {
			return false
		}
	}

	return true
}
