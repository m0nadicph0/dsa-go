package nqueens

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want []Board
	}{
		{
			name: "4x4 board",
			n:    4,
			want: []Board{
				{
					{false, true, false, false},
					{false, false, false, true},
					{true, false, false, false},
					{false, false, true, false},
				},
				{
					{false, false, true, false},
					{true, false, false, false},
					{false, false, false, true},
					{false, true, false, false},
				},
			},
		},
		{
			name: "2x2 board",
			n:    2,
			want: []Board{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveNQueens(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
