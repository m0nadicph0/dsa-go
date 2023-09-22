package lcs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {

	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "Empty inputs",
			text1: "",
			text2: "",
			want:  0,
		},
		{
			name:  "Same Inputs",
			text1: "abcde",
			text2: "abcde",
			want:  5,
		},
		{
			name:  "No common subsequence",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "One of the input is empty",
			text1: "",
			text2: "def",
			want:  0,
		},
		{
			name:  "large texts",
			text1: "awbwcwdwewfw",
			text2: "abcdef",
			want:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestCommonSubsequence(tt.text1, tt.text2)
			assert.Equal(t, tt.want, got)
		})
	}
}
