package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexMatch(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		pattern string
		want    bool
	}{
		{
			name:    "Case 1",
			input:   "aa",
			pattern: "a",
			want:    false,
		},
		{
			name:    "Case 2",
			input:   "aa",
			pattern: "a*",
			want:    true,
		},
		{
			name:    "Case 3",
			input:   "aaa",
			pattern: "a*",
			want:    true,
		},
		{
			name:    "Case 4",
			input:   "aab",
			pattern: "c*a*b",
			want:    true,
		},
		{
			name:    "Case 5",
			input:   "ab",
			pattern: ".*",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RegexMatch(tt.input, tt.pattern), "RegexMatch(%v, %v)", tt.input, tt.pattern)
		})
	}
}
