package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDigits(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "single digit number",
			n:    1,
			want: 1,
		},
		{
			name: "double digit number",
			n:    22,
			want: 2,
		},
		{
			name: "large number",
			n:    1_000_000_000,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountDigits(tt.n), "CountDigits(%v)", tt.n)
		})
	}
}

func TestDigitSum(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "single digit number",
			n:    8,
			want: 8,
		},
		{
			name: "double digit number",
			n:    11,
			want: 2,
		},
		{
			name: "three digit number",
			n:    123,
			want: 6,
		},
		{
			name: "large number",
			n:    1_234_567_890,
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DigitSum(tt.n), "DigitSum(%v)", tt.n)
		})
	}
}

func TestMultiply(t *testing.T) {

	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "Case 1",
			m:    5,
			n:    5,
			want: 25,
		},
		{
			name: "Case 2",
			m:    5,
			n:    25,
			want: 125,
		},
		{
			name: "Case 3",
			m:    25,
			n:    25,
			want: 625,
		},
		{
			name: "Case 4",
			m:    1_000,
			n:    1_000,
			want: 1_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Multiply(tt.m, tt.n), "Multiply(%v, %v)", tt.m, tt.n)
		})
	}
}
