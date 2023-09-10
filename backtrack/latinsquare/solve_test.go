package latinsquare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveLatinSquare(t *testing.T) {

	tests := []struct {
		name string
		grid [][]int
		want bool
	}{
		{
			name: "Case 1",
			grid: NewGrid(4),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveLatinSquare(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
