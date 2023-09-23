package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1",
			nums: []int{1, 2, 3},
			want: 1,
		},
		{
			name: "Case 2",
			nums: []int{4, 6, 5, 2, 4, 3, 2, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.nums...)
			assert.Equal(t, tt.want, got)
		})
	}
}
