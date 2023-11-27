package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		w    int
		want []int
	}{
		{
			name: "Case 1",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			w:    3,
			want: []int{3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Case 2",
			nums: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			w:    4,
			want: []int{3, 3, 3, 3, 3, 3, 3},
		},
		{
			name: "Case 3",
			nums: []int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67},
			w:    3,
			want: []int{10, 9, 23, 23, 34, 56, 67, 67, 67, -1, -2, 9, 10, 34, 67},
		},
		{
			name: "Case 4",
			nums: []int{4, 5, 6, 1, 2, 3},
			w:    1,
			want: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name: "Case 5",
			nums: []int{9, 5, 3, 1, 6, 3},
			w:    2,
			want: []int{9, 5, 3, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxSlidingWindow(tt.nums, tt.w), "MaxSlidingWindow(%v, %v)", tt.nums, tt.w)
		})
	}
}
