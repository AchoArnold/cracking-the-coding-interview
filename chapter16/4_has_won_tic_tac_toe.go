package chapter16

// HasWonTicTacToe checks if a tic tac toe board contains a winner
func HasWonTicTacToe(board [3][3]int) int {
	for i := 0; i < 3; i++ {
		// check rows
		if hasTicTacToeWinner(board[i][0], board[i][1], board[i][2]) {
			return board[i][0]
		}

		// check columns
		if hasTicTacToeWinner(board[0][i], board[1][i], board[2][i]) {
			return board[0][i]
		}

		// check diagonal
		if hasTicTacToeWinner(board[0][0], board[1][1], board[2][2]) {
			return board[0][0]
		}

		if hasTicTacToeWinner(board[0][2], board[1][1], board[2][0]) {
			return board[0][2]
		}
	}

	return 0
}

func hasTicTacToeWinner(p1, p2, p3 int) bool {
	if p1 == 0 {
		return false
	}

	return p1 == p2 && p2 == p3
}
