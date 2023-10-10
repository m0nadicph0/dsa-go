package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {

	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Case 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Case 2",
			height: []int{1, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxArea(tt.height), "MaxArea(%v)", tt.height)
		})
	}
}
