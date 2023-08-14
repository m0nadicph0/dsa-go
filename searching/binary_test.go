package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		key   int
		want  int
	}{
		{
			name:  "Case 1",
			input: []int{1, 2, 4, 7, 8, 9, 25},
			key:   9,
			want:  5,
		},
		{
			name:  "Case 2",
			input: []int{1},
			key:   1,
			want:  0,
		},
		{
			name:  "Case 3",
			input: []int{1, 2},
			key:   2,
			want:  1,
		},
		{
			name:  "Case 4",
			input: []int{1, 2},
			key:   3,
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BinarySearch(tt.input, tt.key), "BinarySearch(%v, %v)", tt.input, tt.key)
		})
	}
}

func TestBinarySearchRange(t *testing.T) {

	tests := []struct {
		name         string
		input        []int
		key          int
		wantFirstPos int
		wantLastPos  int
	}{
		{
			name:         "Case 1",
			input:        []int{1, 2, 3, 3, 3, 3, 4, 6, 8},
			key:          3,
			wantFirstPos: 2,
			wantLastPos:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearchRange(tt.input, tt.key)
			assert.Equalf(t, tt.wantFirstPos, got, "BinarySearchRange(%v, %v)", tt.input, tt.key)
			assert.Equalf(t, tt.wantLastPos, got1, "BinarySearchRange(%v, %v)", tt.input, tt.key)
		})
	}
}
