package algorithms

import (
	"fmt"

	"github.com/palantir/stacktrace"
)

// SudokuSolver solves a sudoku using backtracking
type SudokuSolver struct {
}

var (
	// ErrSudokuCannotBeSolved is thrown when a sudoku cannot be solved
	ErrSudokuCannotBeSolved = stacktrace.NewError("sudoku cannot be solved")
)

// Solve solves a sudoku
func (sudoku SudokuSolver) Solve(board [9][9]int) (result [9][9]int, err error) {
	newBoard := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			newBoard[i][j] = board[i][j]
		}
	}

	return sudoku.backtrack(&newBoard)
}

// IsBoardValid checks if a sudoku board is valid
func (sudoku SudokuSolver) IsBoardValid(board [9][9]int) bool {
	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if sudoku.hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if sudoku.hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if sudoku.hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func (sudoku SudokuSolver) hasEmptyCell(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func (sudoku SudokuSolver) hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

// PrintBoard prints a sudoku board
func (sudoku SudokuSolver) PrintBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func (sudoku SudokuSolver) backtrack(board *[9][9]int) (result [9][9]int, err error) {
	if !sudoku.hasEmptyCell(*board) {
		return *board, nil
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 1; candidate <= 9; candidate++ {
					board[i][j] = candidate
					if sudoku.IsBoardValid(*board) {
						result, err := sudoku.backtrack(board)
						if err == nil {
							return result, nil
						}
					}

					board[i][j] = 0
				}
				return result, ErrSudokuCannotBeSolved
			}
		}
	}
	return result, ErrSudokuCannotBeSolved
}
