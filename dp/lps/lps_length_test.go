package lps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubsequence(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Case 1",
			input: "bbbab",
			want:  4,
		},
		{
			name:  "Case 2",
			input: "cbbd",
			want:  2,
		},
		{
			name:  "Case 3",
			input: "a",
			want:  1,
		},
		{
			name:  "Case 4",
			input: "bananas",
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindromicSubsequence(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
