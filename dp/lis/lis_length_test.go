package lis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestIncSubsequence(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Case 1",
			input: []int{10, 22, 9, 33, 21, 50, 41, 60, 80},
			want:  6,
		},
		{
			name:  "Case 2",
			input: []int{3, 4, 5, 10},
			want:  4,
		},
		{
			name:  "Case 3",
			input: []int{50, 3, 10, 7, 40, 80},
			want:  4,
		},
		{
			name:  "Case 4",
			input: []int{3, 2, 6, 4, 5},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestIncSubsequence(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
