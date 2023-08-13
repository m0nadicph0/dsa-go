package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {

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
		{
			name:  "Case 2",
			input: []int{2, 1, 5, 0, 17, 7, 6},
			want:  []int{0, 1, 2, 5, 6, 7, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.input)
			assert.Equal(t, tt.want, tt.input, "QuickSort")
		})
	}
}
