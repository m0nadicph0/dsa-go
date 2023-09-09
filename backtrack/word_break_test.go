package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreak(t *testing.T) {

	tests := []struct {
		name  string
		input string
		words []string
		want  bool
	}{
		{
			name:  "Case 1",
			input: "leetcode",
			words: []string{"leet", "code"},
			want:  true,
		},
		{
			name:  "Case 2",
			input: "applepenapple",
			words: []string{"apple", "pen"},
			want:  true,
		},
		{
			name:  "Case 3",
			input: "catsandog",
			words: []string{"cats", "dog", "sand", "and", "cat"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WordBreak(tt.input, tt.words), "WordBreak(%v, %v)", tt.input, tt.words)
		})
	}
}
