package knight

import "fmt"

type Board [][]int

const EmptyCellValue = -1

func NewBoard(size int) Board {
	board := make([][]int, size)
	for x := 0; x < size; x++ {
		board[x] = make([]int, size)
		for y := 0; y < size; y++ {
			board[x][y] = EmptyCellValue
		}
	}
	return board
}
func (b Board) Rows() int {
	return len(b)
}

func (b Board) Cols() int {
	return len(b[0])
}

func (b Board) IsEmpty(row, col int) bool {
	return b[row][col] == EmptyCellValue
}

func (b Board) IsSafe(row, col int) bool {
	return row >= 0 && row < b.Rows() && col >= 0 && col < b.Cols() && b.IsEmpty(row, col)
}

func (b Board) Print() {
	for x := 0; x < b.Rows(); x++ {
		for y := 0; y < b.Cols(); y++ {
			fmt.Printf("%2d ", b[x][y])
		}
		fmt.Println()
	}
}

func (b Board) Put(row int, col int, value int) {
	b[row][col] = value
}

func (b Board) Cells() int {
	return b.Rows() * b.Cols()
}

func (b Board) Clear(row int, col int) {
	b[row][col] = EmptyCellValue
}
