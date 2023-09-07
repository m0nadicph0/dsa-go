package knight

type Move struct {
	X int
	Y int
}

func KnightsTour(board Board, startX, startY int) bool {

	// moves define the next move of the Knight
	moves := []Move{
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}

	// Since the Knight is initially at the first block
	board.Put(startX, startY, 0)

	// Start from 0,0 and explore all tours using KnightsTourHelper()
	return KnightsTourHelper(board, startX, startY, 1, moves)
}

func KnightsTourHelper(board Board, startX, startY, moveNo int, moves []Move) bool {
	if moveNo == board.Cells() {
		return true
	}

	// Try all next moves from the current coordinate startX, startY
	for k := 0; k < 8; k++ {
		nextX := startX + moves[k].X
		nextY := startY + moves[k].Y
		if board.IsSafe(nextX, nextY) {
			board.Put(nextX, nextY, moveNo)
			if KnightsTourHelper(board, nextX, nextY, moveNo+1, moves) {
				return true
			}
			board.Clear(nextX, nextY) // backtracking
		}
	}

	return false
}
