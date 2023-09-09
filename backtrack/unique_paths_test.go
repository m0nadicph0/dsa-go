package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {

	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "Case 1x1",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "Case 3x7",
			m:    3,
			n:    7,
			want: 28,
		}, {
			name: "Case 4x4",
			m:    4,
			n:    4,
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UniquePaths(tt.m, tt.n), "UniquePaths(%v, %v)", tt.m, tt.n)
		})
	}
}
