package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Case 1",
			input: "23",
			want:  []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:  "Case 2",
			input: "2",
			want:  []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LetterCombinations(tt.input), "LetterCombinations(%v)", tt.input)
		})
	}
}
