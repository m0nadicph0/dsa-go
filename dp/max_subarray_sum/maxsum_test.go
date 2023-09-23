package max_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarraySum(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "Case 2",
			nums: []int{1, 2, 3, 4},
			want: 10,
		},
		{
			name: "Case 3",
			nums: []int{1, 2, 3, 4, -2, 5},
			want: 13,
		},
		{
			name: "Case 4",
			nums: []int{-1, -2, -3, -4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaximumSubarraySum(tt.nums))
		})
	}
}
