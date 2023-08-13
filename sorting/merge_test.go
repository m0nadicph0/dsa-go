package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Base case 0",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Base case 1",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "Case 0",
			input: []int{5, 4},
			want:  []int{4, 5},
		},
		{
			name:  "Case 1",
			input: []int{5, 4, 3, 2, 1, 0},
			want:  []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.input)
			assert.Equal(t, tt.want, tt.input, "MergeSort")
		})
	}
}
