package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		key   int
		want  int
	}{
		{
			name:  "case 1",
			input: []int{1, 2, 3, 4},
			key:   3,
			want:  2,
		},
		{
			name:  "case 2",
			input: []int{1, 2, 3, 4},
			key:   5,
			want:  -1,
		},
		{
			name:  "case 3",
			input: []int{},
			key:   5,
			want:  -1,
		},
		{
			name:  "case 4",
			input: []int{1, 2, 2, 2, 1, 1, 1, 3, 4, 2, 4, 3},
			key:   3,
			want:  7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LinearSearch(tt.input, tt.key)
			assert.Equal(t, tt.want, got, "LinearSearch(%v,%d) = %d", tt.input, tt.key, got)
		})
	}
}
