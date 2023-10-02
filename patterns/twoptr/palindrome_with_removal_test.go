package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPalindromeWithOneRemoval(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Case 1",
			s:    "madame",
			want: true,
		},
		{
			name: "Case 2",
			s:    "dead",
			want: true,
		},
		{
			name: "Case 3",
			s:    "abca",
			want: true,
		},
		{
			name: "Case 4",
			s:    "tebbem",
			want: false,
		},
		{
			name: "Case 5",
			s:    "eeccccbebaeeabebccceea",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ValidPalindromeWithOneRemoval(tt.s), "ValidPalindromeWithOneRemoval(%v)", tt.s)
		})
	}
}
