package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateExpressions(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		target int
		want   []string
	}{
		{
			name:   "Case 1",
			input:  "123",
			target: 6,
			want:   []string{"1*2*3", "1+2+3"},
		},
		{
			name:   "Case 2",
			input:  "232",
			target: 8,
			want:   []string{"2*3+2", "2+3*2"},
		},
		{
			name:   "Case 3",
			input:  "3456237490",
			target: 9191,
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, GenerateExpressions(tt.input, tt.target), "GenerateExpressions(%v, %v)", tt.input, tt.target)
		})
	}
}
