package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetSum(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		target int
		want   bool
	}{
		{
			name:   "Case 1",
			input:  []int{3, 34, 4, 12, 5, 2},
			target: 9,
			want:   true,
		},
		{
			name:   "Case 2",
			input:  []int{3, 34, 4, 12, 5, 2},
			target: 30,
			want:   false,
		},
		{
			name:   "Case 3",
			input:  []int{2, 3, 5, 6, 8, 10},
			target: 10,
			want:   true,
		},
		{
			name:   "Case 4",
			input:  []int{2, 3},
			target: 60,
			want:   false,
		},
		{
			name:   "Case 5",
			input:  []int{},
			target: 1,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SubsetSum(tt.input, tt.target), "SubsetSum(%v, %v)", tt.input, tt.target)
		})
	}
}
