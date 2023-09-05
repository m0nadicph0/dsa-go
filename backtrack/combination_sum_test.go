package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Case 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Case 2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Case 3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CombinationSum(tt.candidates, tt.target), "CombinationSum(%v, %v)", tt.candidates, tt.target)
		})
	}
}
