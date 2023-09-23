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
			input: "BBABCBCAB",
			want:  7,
		},
		{
			name:  "Case 5",
			input: "bebeeed",
			want:  4,
		},
		{
			name:  "Case 6",
			input: "abcde",
			want:  1,
		},
		{
			name:  "Case 7",
			input: "abbaab",
			want:  4,
		},
		{
			name:  "Case 8",
			input: "aa",
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindromicSubsequence(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
