package sudoku

import "fmt"

type Grid [][]int

func NewGrid(config [][]int) Grid {
	grid := make([][]int, len(config))
	for i := 0; i < len(config); i++ {
		row := make([]int, len(config[0]))
		for j := 0; j < len(config[0]); j++ {
			row[j] = config[i][j]
		}
		grid[i] = row
	}
	return grid
}

func (g Grid) Print() {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			fmt.Printf("%d ", g[i][j])
		}
		fmt.Println()
	}
}

func (g Grid) IsValid(row, col, value int) bool {
	// Check the row
	for i := 0; i < len(g); i++ {
		if g[row][i] == value {
			return false
		}
	}

	// Check the column
	for i := 0; i < len(g[0]); i++ {
		if g[i][col] == value {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g[i+startRow][j+startCol] == value {
				return false
			}
		}
	}

	return true
}

func (g Grid) Rows() int {
	return len(g)
}

func (g Grid) Cols() int {
	return len(g[0])
}

func (g Grid) Clear(row, col int) {
	g[row][col] = 0
}

func (g Grid) Put(row, col, value int) {
	g[row][col] = value
}

func (g Grid) IsEmpty(row, col int) bool {
	return g[row][col] == 0
}
