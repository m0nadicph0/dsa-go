package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "Case 1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "Case 2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "Case 3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TwoSum(tt.numbers, tt.target), "TwoSum(%v, %v)", tt.numbers, tt.target)
		})
	}
}
