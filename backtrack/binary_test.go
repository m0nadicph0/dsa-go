package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryStrings(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Case 0",
			n:    0,
			want: []string{""},
		},
		{
			name: "Case 1",
			n:    1,
			want: []string{"0", "1"},
		},
		{
			name: "Case 2",
			n:    2,
			want: []string{"00", "01", "10", "11"},
		},
		{
			name: "Case 3",
			n:    3,
			want: []string{"000", "001", "010", "011", "100", "101", "110", "111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BinaryStrings(tt.n), "BinaryStrings(%v)", tt.n)
		})
	}
}
