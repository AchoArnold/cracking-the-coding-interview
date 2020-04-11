package chapter10

import (
	"math"
)

// MatrixCoordinate represents teh coordinates of a matrix
type MatrixCoordinate struct {
	row    int
	column int
}

// NewMatrixCoordinate creates a new coordinate
func NewMatrixCoordinate(row, col int) MatrixCoordinate {
	return MatrixCoordinate{
		row:    row,
		column: col,
	}
}

func (mc MatrixCoordinate) inbounds(matrix [][]int) bool {
	return mc.row >= 0 && mc.column >= 0 && mc.row < len(matrix) && mc.column < len(matrix[0])
}

func (mc MatrixCoordinate) clone() MatrixCoordinate {
	return MatrixCoordinate{
		row:    mc.row,
		column: mc.column,
	}
}

func (mc *MatrixCoordinate) setToAverage(min, max MatrixCoordinate) {
	mc.row = (min.row + max.row) / 2
	mc.column = (min.column + max.column) / 2
}

func (mc MatrixCoordinate) isBefore(end MatrixCoordinate) bool {
	return mc.row <= end.row && mc.column <= end.column
}

// FindElementInSortedMatrix finds x in a matrix that is sorted by both row and column
func FindElementInSortedMatrix(matrix [][]int, x int) *MatrixCoordinate {
	origin := NewMatrixCoordinate(0, 0)
	dest := NewMatrixCoordinate(len(matrix)-1, len(matrix[0])-1)

	return findElementInSortedMatrix(matrix, origin, dest, x)
}

func findElementInSortedMatrix(matrix [][]int, origin, dest MatrixCoordinate, x int) *MatrixCoordinate {
	if !origin.inbounds(matrix) || !dest.inbounds(matrix) {
		return nil
	}

	if matrix[origin.row][origin.column] == x {
		return &origin
	}

	// Set start to start of diagonal and end tot he end of the diagonal. Since the grid may not be square. The end of
	// the grid may not be square, the end of the diagonal may not be equal to the destination
	start := origin.clone()
	diagDist := int(math.Min(float64(dest.row-origin.row), float64(dest.column-origin.column)))
	end := NewMatrixCoordinate(start.row+diagDist, start.column+diagDist)
	p := NewMatrixCoordinate(0, 0)

	// Do binary search on the diagonal, looking for the first element > x */
	for start.isBefore(end) {
		p.setToAverage(start, end)
		if x > matrix[p.row][p.column] {
			start.row = p.row + 1
			start.column = p.column + 1
		} else {
			end.row = p.row - 1
			end.column = p.column - 1
		}
	}

	// Split the grid into quadrants. Search the bottom left and the top right
	return partitionSearch(matrix, origin, dest, start, x)
}

func partitionSearch(matrix [][]int, origin, dest, pivot MatrixCoordinate, x int) *MatrixCoordinate {
	lowerLeftOrigin := NewMatrixCoordinate(pivot.row, origin.column)
	lowerLeftDest := NewMatrixCoordinate(dest.row, pivot.column-1)
	upperRightOrigin := NewMatrixCoordinate(origin.row, pivot.column)
	upperRightDest := NewMatrixCoordinate(pivot.row-1, dest.column)

	lowerLeft := findElementInSortedMatrix(matrix, lowerLeftOrigin, lowerLeftDest, x)
	if lowerLeft == nil {
		return findElementInSortedMatrix(matrix, upperRightOrigin, upperRightDest, x)
	}

	return lowerLeft
}
