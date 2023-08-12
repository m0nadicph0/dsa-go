package recursion

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "non empty unsorted array",
			input:    []int{3, 4, 3, 2, 1, 0},
			expected: "343210",
		},
		{
			name:     "non empty sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: "12345",
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := bytes.NewBuffer(make([]byte, 0))
			Print(out, tc.input)
			assert.Equal(t, tc.expected, out.String(), "they should be equal")
		})
	}

}

func TestReversePrint(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "non empty unsorted array",
			input:    []int{3, 4, 3, 2, 1, 0},
			expected: "012343",
		},
		{
			name:     "non empty sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: "54321",
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := bytes.NewBuffer(make([]byte, 0))
			ReversePrint(out, tc.input)
			assert.Equal(t, tc.expected, out.String(), "they should be equal")
		})
	}

}

func TestReversePrintAlt(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "non empty unsorted array",
			input:    []int{3, 4, 3, 2, 1, 0},
			expected: "012343",
		},
		{
			name:     "non empty sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: "54321",
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := bytes.NewBuffer(make([]byte, 0))
			ReversePrintAlt(out, tc.input)
			assert.Equal(t, tc.expected, out.String(), "they should be equal")
		})
	}

}

func TestLength(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "non empty unsorted array",
			input:    []int{3, 4, 3, 2, 1, 0},
			expected: 6,
		},
		{
			name:     "non empty sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Length(tc.input)
			assert.Equal(t, tc.expected, actual, "they should be equal")
		})
	}

}

func TestReplace(t *testing.T) {

	testCases := []struct {
		name        string
		input       []int
		target      int
		replacement int
		expected    []int
	}{
		{
			name:        "non empty unsorted array",
			input:       []int{3, 4, 3, 2, 1, 0},
			target:      3,
			replacement: 1,
			expected:    []int{1, 4, 1, 2, 1, 0},
		},
		{
			name:        "non empty array with target element",
			input:       []int{1, 2, 3, 4, 5},
			target:      2,
			replacement: 0,
			expected:    []int{1, 0, 3, 4, 5},
		},
		{
			name:        "non empty array without target element",
			input:       []int{1, 2, 3, 4, 5},
			target:      10,
			replacement: 0,
			expected:    []int{1, 2, 3, 4, 5},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Replace(tc.input, tc.target, tc.replacement)
			assert.Equal(t, tc.expected, tc.input, "they should be equal")
		})
	}

}

func TestIsSorted(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "base case 1",
			args: []int{},
			want: true,
		},
		{
			name: "base case 2",
			args: []int{1},
			want: true,
		},
		{
			name: "sorted array",
			args: []int{1, 2, 3, 4},
			want: true,
		},
		{
			name: "unsorted array",
			args: []int{4, 1, 2, 3},
			want: false,
		},
		{
			name: "sorted array in reverse order",
			args: []int{4, 3, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsSorted(tt.args), "IsSorted(%v)", tt.args)
		})
	}
}
