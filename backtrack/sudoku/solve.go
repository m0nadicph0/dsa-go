package sudoku

func SolveSudoku(grid Grid) bool {
	for row := 0; row < grid.Rows(); row++ {
		for col := 0; col < grid.Cols(); col++ {
			// If the cell is empty (0), try filling it with numbers 1-9
			if grid.IsEmpty(row, col) {
				for num := 1; num <= 9; num++ {
					if grid.IsValid(row, col, num) {
						// Try placing 'num' in the cell
						grid.Put(row, col, num)

						// Recursively solve the rest of the puzzle
						if SolveSudoku(grid) {
							return true
						}

						// If placing 'num' didn't lead to a solution, backtrack
						grid.Clear(row, col)
					}
				}
				// If no number can be placed, the puzzle is unsolvable at this point
				return false
			}
		}
	}
	// If all cells are filled, the puzzle is solved
	return true
}
