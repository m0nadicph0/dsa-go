package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Case 0",
			input: []int{3},
			want:  []int{3},
		},
		{
			name:  "Case 1",
			input: []int{3, 2, 1},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Case 2",
			input: []int{3, 2, 1, 4, 6, 5},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "Case 3",
			input: []int{4, 3, 2, 1},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Case 4",
			input: []int{1, 2, 2, 1},
			want:  []int{1, 1, 2, 2},
		},
		{
			name:  "Case 5",
			input: []int{100, 99, 89, 78, 67, 55, 44, 33, 22, 10, 5, 4, 3},
			want:  []int{3, 4, 5, 10, 22, 33, 44, 55, 67, 78, 89, 99, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.input)
			assert.Equal(t, tt.want, tt.input, "BubbleSort")
		})
	}
}
