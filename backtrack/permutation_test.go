package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutations(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{
			name:  "empty input",
			input: []int{},
			want:  [][]int{},
		},
		{
			name:  "single element input",
			input: []int{1},
			want:  [][]int{{1}},
		},
		{
			name:  "two element input",
			input: []int{1, 2},
			want:  [][]int{{1, 2}, {2, 1}},
		},
		{
			name:  "three element input",
			input: []int{1, 2, 3},
			want:  [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutations(tt.input)
			assert.ElementsMatch(t, tt.want, got, "Permutations(%v)=%v", tt.input, got)
		})
	}
}
