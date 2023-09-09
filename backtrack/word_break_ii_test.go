package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakWordToSentences(t *testing.T) {

	tests := []struct {
		name  string
		input string
		words []string
		want  []string
	}{
		{
			name:  "Case 1",
			input: "catsanddog",
			words: []string{"cat", "cats", "and", "sand", "dog"},
			want:  []string{"cat sand dog", "cats and dog"},
		},
		{
			name:  "Case 2",
			input: "pineapplepenapple",
			words: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:  []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, BreakWordToSentences(tt.input, tt.words), "BreakWordToSentences(%v, %v)", tt.input, tt.words)
		})
	}
}
