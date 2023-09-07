package nsudoku

func solveSudoku(grid Grid, fn func(Grid)) {
	for row := 0; row < grid.Rows(); row++ {
		for col := 0; col < grid.Cols(); col++ {
			if grid.IsEmpty(row, col) {
				for value := 0; value < 10; value++ {
					if grid.IsValid(row, col, value) {
						grid.Put(row, col, value)
						solveSudoku(grid, fn)
						grid.Clear(row, col)
					}
				}
				return
			}
		}
	}
	fn(grid)
}

func CountSolutions(cfg [][]int) int {
	count := 0
	solveSudoku(NewGrid(cfg), func(grid Grid) {
		count += 1
	})

	return count
}
