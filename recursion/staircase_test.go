package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountWaysToClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "base case",
			n:    0,
			want: 1,
		},
		{
			name: "case 1",
			n:    1,
			want: 1,
		},
		{
			name: "case 2",
			n:    2,
			want: 2,
		},
		{
			name: "case 3",
			n:    3,
			want: 4,
		},
		{
			name: "case 4",
			n:    4,
			want: 7,
		},
		{
			name: "case 5",
			n:    5,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountWaysToClimbStairs(tt.n), "CountWaysToClimbStairs(%v)", tt.n)
		})
	}
}
