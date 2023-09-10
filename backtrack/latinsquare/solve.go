package latinsquare

func solveLatinSquareHelper(grid Grid, row, col int) bool {
	if row == grid.Rows() {
		return true
	}

	for num := 1; num <= grid.Rows(); num++ {
		if grid.IsSafe(row, col, num) {
			grid[row][col] = num

			nextRow, nextCol := grid.NextCell(row, col)

			if solveLatinSquareHelper(grid, nextRow, nextCol) {
				return true
			}
			grid[row][col] = 0 // Backtrack
		}
	}
	return false
}
func SolveLatinSquare(grid [][]int) bool {
	return solveLatinSquareHelper(grid, 0, 0)
}
