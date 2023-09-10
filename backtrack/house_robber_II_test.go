package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobII(t *testing.T) {

	tests := []struct {
		name   string
		wealth []int
		want   int
	}{
		{
			name:   "Case 1",
			wealth: []int{2, 3, 2},
			want:   3,
		},
		{
			name:   "Case 2",
			wealth: []int{1, 2, 3, 1},
			want:   4,
		},
		{
			name:   "Case 3",
			wealth: []int{1, 2, 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RobII(tt.wealth), "RobII(%v)", tt.wealth)
		})
	}
}
