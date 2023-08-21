package nqueens

type visitFn func(b Board)

func SolveNQueens(n int) []Board {
	result := make([]Board, 0)
	board := NewBoard(n)
	forEachSolution(board, 0, func(b Board) {
		result = append(result, b.Copy())
	})
	return result
}

func forEachSolution(board Board, row int, fn visitFn) {
	if row == board.Rows() {
		fn(board)
		return
	}

	for col := 0; col < board.Cols(); col++ {
		if board.IsValid(row, col) {
			board.PlaceQueen(row, col)
			forEachSolution(board, row+1, fn)
			board.Clear(row, col)
		}
	}
}
