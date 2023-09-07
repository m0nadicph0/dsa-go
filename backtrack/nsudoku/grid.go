package nsudoku

type Grid [][]int

func (g Grid) Rows() int {
	return len(g)
}

func (g Grid) Cols() int {
	return len(g[0])
}

func (g Grid) IsEmpty(row int, col int) bool {
	return g[row][col] == 0
}

func (g Grid) IsValid(row int, col int, value int) bool {
	for i := 0; i < g.Cols(); i++ {
		if g[row][i] == value {
			return false
		}
	}

	for i := 0; i < g.Rows(); i++ {
		if g[i][col] == value {
			return false
		}
	}
	startR := row - row%3
	startC := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g[startR+i][startC+j] == value {
				return false
			}
		}
	}
	return true
}

func (g Grid) Put(row int, col int, value int) {
	g[row][col] = value
}

func (g Grid) Clear(row int, col int) {
	g[row][col] = 0
}

func NewGrid(cfg [][]int) Grid {
	grid := make(Grid, len(cfg))
	for row := 0; row < len(cfg); row++ {
		grid[row] = make([]int, len(cfg[0]))
		copy(grid[row], cfg[row])
	}
	return grid
}
