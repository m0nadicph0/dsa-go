package latinsquare

import "fmt"

type Grid [][]int

func NewGrid(size int) Grid {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}
	return grid
}

func (g Grid) Rows() int {
	return len(g)
}

func (g Grid) Cols() int {
	return len(g[0])
}

func (g Grid) NextCell(row, col int) (int, int) {
	col++
	if col == g.Cols() {
		col = 0
		row++
	}
	return row, col
}

func (g Grid) Print() {
	for i := 0; i < g.Rows(); i++ {
		for j := 0; j < g.Cols(); j++ {
			fmt.Printf("%d ", g[i][j])
		}
		fmt.Println()
	}
}

func (g Grid) IsSafe(row, col, value int) bool {
	for i := 0; i < g.Rows(); i++ {
		if g[row][i] == value || g[i][col] == value {
			return false
		}
	}
	return true
}
