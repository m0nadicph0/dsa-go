package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsOfSize(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		size  int
		want  [][]int
	}{
		{
			name:  "Case 1",
			input: []int{1, 2, 3},
			size:  2,
			want:  [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name:  "Case 2",
			input: []int{1, 2, 3},
			size:  1,
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "Case 3",
			input: []int{1, 2, 3},
			size:  3,
			want:  [][]int{{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SubsetsOfSize(tt.input, tt.size), "SubsetsOfSize(%v, %v)", tt.input, tt.size)
		})
	}
}
