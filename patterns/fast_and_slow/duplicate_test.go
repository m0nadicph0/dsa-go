package fast_and_slow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1",
			nums: []int{3, 4, 4, 4, 2},
			want: 4,
		},
		{
			name: "Case 2",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "Case 3",
			nums: []int{1, 3, 6, 2, 7, 3, 5, 4},
			want: 3,
		},
		{
			name: "Case 4",
			nums: []int{1, 2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindDuplicate(tt.nums), "FindDuplicate(%v)", tt.nums)
		})
	}
}
