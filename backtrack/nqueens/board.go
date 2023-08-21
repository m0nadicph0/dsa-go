package nqueens

import (
	"bytes"
	"fmt"
)

type Board [][]bool

func NewBoard(n int) Board {
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		row := make([]bool, n)
		for j := 0; j < n; j++ {
			row[j] = false
		}
		board[i] = row
	}
	return board
}

func (b Board) HasQueen(row, col int) bool {
	return b[row][col]
}

func (b Board) PlaceQueen(row, col int) {
	b[row][col] = true
}

func (b Board) Clear(row, col int) {
	b[row][col] = false
}

func (b Board) IsValid(row, col int) bool {
	for i := 0; i < row; i++ {
		if b.HasQueen(i, col) {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.HasQueen(i, j) {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < b.Cols(); i, j = i-1, j+1 {
		if b.HasQueen(i, j) {
			return false
		}
	}

	return true
}

func (b Board) Rows() int {
	return len(b)
}

func (b Board) Cols() int {
	return len(b[0])
}

func (b Board) Copy() Board {
	tmpBoard := make([][]bool, len(b))
	for i := 0; i < len(b); i++ {
		row := make([]bool, len(b[i]))
		copy(row, b[i])
		tmpBoard[i] = row
	}
	return tmpBoard
}

func (b Board) String() string {
	buff := new(bytes.Buffer)

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] {
				fmt.Fprint(buff, "Q")
			} else {
				fmt.Fprint(buff, ".")
			}
		}
		fmt.Fprintln(buff, "")
	}
	return buff.String()
}
