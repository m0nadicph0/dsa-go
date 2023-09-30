package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Case 1",
			input: "racecar",
			want:  true,
		},
		{
			name:  "Case 2",
			input: "abab",
			want:  false,
		},
		{
			name:  "Case 3",
			input: "bb",
			want:  true,
		},
		{
			name:  "Case 4",
			input: "a",
			want:  true,
		},
		{
			name:  "Case 5",
			input: "ablewasiereisawelba",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsPalindrome(tt.input))
		})
	}
}
