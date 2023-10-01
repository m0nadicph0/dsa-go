package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		colors []int
		want   []int
	}{
		{
			name:   "Case 1",
			colors: []int{0, 1, 0},
			want:   []int{0, 0, 1},
		},
		{
			name:   "Case 2",
			colors: []int{1},
			want:   []int{1},
		},
		{
			name:   "Case 3",
			colors: []int{2, 2},
			want:   []int{2, 2},
		},
		{
			name:   "Case 4",
			colors: []int{1, 1, 0, 2},
			want:   []int{0, 1, 1, 2},
		},
		{
			name:   "Case 5",
			colors: []int{2, 1, 1, 0, 0},
			want:   []int{0, 0, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SortColors(tt.colors), "SortColors(%v)", tt.colors)
		})
	}
}
