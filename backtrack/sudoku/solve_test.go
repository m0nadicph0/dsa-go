package sudoku

import "testing"

func TestSolveSudoku(t *testing.T) {

	tests := []struct {
		name   string
		config [][]int
		want   bool
	}{
		{
			name: "Case 1",
			config: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			want: true,
		},
		{
			name: "Case 2",
			config: [][]int{
				{5, 1, 0, 0, 9, 0, 7, 0, 0},
				{6, 0, 9, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 7, 0, 0, 0, 8},
				{0, 3, 0, 0, 0, 0, 5, 0, 0},
				{0, 0, 0, 5, 0, 8, 0, 0, 0},
				{0, 0, 4, 0, 0, 0, 0, 6, 0},
				{2, 0, 0, 0, 8, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 9, 0, 7},
				{0, 0, 6, 0, 5, 0, 0, 4, 3},
			},
			want: true,
		},
		{
			name: "Case 3",
			config: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid(tt.config)
			if got := SolveSudoku(grid); got != tt.want {
				t.Errorf("SolveSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
