package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Case 1",
			input: "0000",
			want:  []string{"0.0.0.0"},
		},
		{
			name:  "Case 2",
			input: "25525511135",
			want:  []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name:  "Case 3",
			input: "1111",
			want:  []string{"1.1.1.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, IPAddresses(tt.input), "IPAddresses(%v)", tt.input)
		})
	}
}
