package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {

	tests := []struct {
		name   string
		wealth []int
		want   int
	}{
		{
			name:   "Case 1",
			wealth: []int{1, 2, 3, 1},
			want:   4,
		},
		{
			name:   "Case 2",
			wealth: []int{2, 7, 9, 3, 1},
			want:   12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Rob(tt.wealth), "Rob(%v)", tt.wealth)
		})
	}
}
