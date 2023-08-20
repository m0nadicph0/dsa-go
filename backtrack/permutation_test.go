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

func TestPermutationsNK(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		k     int
		want  [][]int
	}{
		{
			name:  "empty input",
			input: []int{},
			k:     0,
			want:  [][]int{{}},
		},
		{
			name:  "single element input",
			input: []int{1},
			k:     1,
			want:  [][]int{{1}},
		},
		{
			name:  "two element input with k=1",
			input: []int{1, 2},
			k:     1,
			want:  [][]int{{1}, {2}},
		},
		{
			name:  "two element input with k=2",
			input: []int{1, 2},
			k:     2,
			want:  [][]int{{1, 2}, {2, 1}},
		},
		{
			name:  "three element input with k=1",
			input: []int{1, 2, 3},
			k:     1,
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "three element input with k=2",
			input: []int{1, 2, 3},
			k:     2,
			want:  [][]int{{1, 2}, {2, 1}, {1, 3}, {3, 1}, {2, 3}, {3, 2}},
		},
		{
			name:  "three element input with k=3",
			input: []int{1, 2, 3},
			k:     3,
			want:  [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, PermutationsNK(tt.input, tt.k), "PermutationsNK(%v, %v)", tt.input, tt.k)
		})
	}
}
