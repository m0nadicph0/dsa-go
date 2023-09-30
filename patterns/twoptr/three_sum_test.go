package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Case 1",
			nums:   []int{1, -1, 0},
			target: -1,
			want:   false,
		},
		{
			name:   "Case 2",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 10,
			want:   true,
		},
		{
			name:   "Case 3",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 21,
			want:   false,
		},
		{
			name:   "Case 4",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			want:   true,
		},
		{
			name:   "Case 5",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ThreeSum(tt.nums, tt.target), "ThreeSum(%v, %v)", tt.nums, tt.target)
		})
	}
}
