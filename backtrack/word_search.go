package backtrack

type WordBoard [][]byte

func (b WordBoard) Rows() int {
	return len(b)
}

func (b WordBoard) Cols() int {
	return len(b[0])
}

func wordSearchHelper(board WordBoard, word string, visited [][]bool, row, col, index int) bool {
	rows := board.Rows()
	cols := board.Cols()

	if index == len(word) {
		return true // Word found
	}

	// Check if the current cell is within bounds and matches the current character
	if row >= 0 && row < rows && col >= 0 && col < cols && board[row][col] == word[index] && !visited[row][col] {
		visited[row][col] = true

		// Explore adjacent cells in four directions
		if wordSearchHelper(board, word, visited, row+1, col, index+1) ||
			wordSearchHelper(board, word, visited, row-1, col, index+1) ||
			wordSearchHelper(board, word, visited, row, col+1, index+1) ||
			wordSearchHelper(board, word, visited, row, col-1, index+1) {
			return true
		}

		visited[row][col] = false // Backtrack
	}

	return false
}
func makeVisitTable(board WordBoard) [][]bool {
	visited := make([][]bool, board.Rows())
	for i := 0; i < board.Rows(); i++ {
		visited[i] = make([]bool, board.Cols())
	}
	return visited
}

func WordSearch(board WordBoard, word string) bool {
	visited := makeVisitTable(board)

	// Iterate through the board to find the starting point
	for i := 0; i < board.Rows(); i++ {
		for j := 0; j < board.Cols(); j++ {
			if board[i][j] == word[0] && wordSearchHelper(board, word, visited, i, j, 0) {
				return true // Word found starting from (i, j)
			}
		}
	}

	return false // Word not found
}
