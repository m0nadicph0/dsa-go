package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinations(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		k     int
		want  [][]int
	}{
		{
			name:  "empty input",
			input: []int{},
			k:     1,
			want:  [][]int{},
		},
		{
			name:  "single element input",
			input: []int{1},
			k:     1,
			want:  [][]int{{1}},
		},
		{
			name:  "3 element input with k=2",
			input: []int{1, 2, 3},
			k:     2,
			want:  [][]int{{1, 2}, {2, 3}, {1, 3}},
		},
		{
			name:  "3 element input with k=3",
			input: []int{1, 2, 3},
			k:     3,
			want:  [][]int{{1, 2, 3}},
		},
		{
			name:  "3 element input with k=1",
			input: []int{1, 2, 3},
			k:     1,
			want:  [][]int{{1}, {2}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combinations(tt.input, tt.k)
			assert.ElementsMatch(t, tt.want, got, "Combinations(%v, %d)", tt.input, tt.k)
		})
	}
}
