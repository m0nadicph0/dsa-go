package decode_ways

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeWays(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "12 could be decoded as AB (1 2) or L (12)",
			input: "12",
			want:  2,
		},
		{
			name:  "123 could be decoded as ABC (1 2 3) or LC (12 3) or AW (1 23)",
			input: "123",
			want:  3,
		},
		{
			name:  "1234 could be decoded as ABCD (1 2 3 4) or LCD (12 3 4) or AWD (1 23 4)",
			input: "1234",
			want:  3,
		},
		{
			name:  "226 could be decoded as BZ (2 26), VF (22 6), or BBF (2 2 6)",
			input: "226",
			want:  3,
		},
		{
			name:  "06 cannot be mapped to F because of the leading zero",
			input: "06",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeWays(tt.input)
			assert.Equal(t, tt.want, got, "DecodeWays(%s)=%d", tt.input, got)
		})
	}
}
