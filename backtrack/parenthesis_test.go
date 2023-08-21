package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParenthesisCombinations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "One pair of parenthesis",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "Two pairs of parenthesis",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "Three pairs of parenthesis",
			n:    3,
			want: []string{"((()))", "(()())", "()(())", "(())()", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, GenerateParenthesisCombinations(tt.n), "GenerateParenthesisCombinations(%v)", tt.n)
		})
	}
}
