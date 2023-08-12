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

func TestSum(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "empty array",
			args: []int{},
			want: 0,
		},
		{
			name: "single element array",
			args: []int{1},
			want: 1,
		},
		{
			name: "sorted array",
			args: []int{1, 2, 3, 4, 5},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Sum(tt.args), "Sum(%v)", tt.args)
		})
	}
}

func TestIsPresent(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		target int
		want   bool
	}{
		{
			name:   "empty array",
			input:  []int{},
			target: 0,
			want:   false,
		},
		{
			name:   "single element array with target",
			input:  []int{1},
			target: 1,
			want:   true,
		},
		{
			name:   "single element array without target",
			input:  []int{1},
			target: 2,
			want:   false,
		},
		{
			name:   "multiple element array with target at start",
			input:  []int{1, 2, 3, 4},
			target: 1,
			want:   true,
		},
		{
			name:   "multiple element array with target at end",
			input:  []int{1, 2, 3, 4},
			target: 4,
			want:   true,
		},
		{
			name:   "multiple element array without target",
			input:  []int{1, 2, 3, 4},
			target: 5,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsPresent(tt.input, tt.target), "IsPresent(%v, %v)", tt.input, tt.target)
		})
	}
}

func TestFirstIndexOf(t *testing.T) {

	tests := []struct {
		name   string
		args   []int
		target int
		want   int
	}{
		{
			name:   "base case",
			args:   []int{},
			target: 1,
			want:   -1,
		},
		{
			name:   "single element array with target",
			args:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "non empty array with target at start",
			args:   []int{1, 0, 20, 10},
			target: 1,
			want:   0,
		},
		{
			name:   "non empty array with target at end",
			args:   []int{1, 0, 20, 10},
			target: 10,
			want:   3,
		},
		{
			name:   "non empty array with repetition",
			args:   []int{1, 0, 0, 1},
			target: 0,
			want:   1,
		},
		{
			name:   "non empty array with repetition at start",
			args:   []int{1, 1, 0, 1},
			target: 1,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FirstIndexOf(tt.args, tt.target), "FirstIndexOf(%v)", tt.args)
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		target int
		want   int
	}{
		{
			name:   "base case",
			args:   []int{},
			target: 1,
			want:   -1,
		},
		{
			name:   "single element array with target",
			args:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "non empty array with target at start",
			args:   []int{1, 0, 20, 10},
			target: 1,
			want:   0,
		},
		{
			name:   "non empty array with target at end",
			args:   []int{1, 0, 20, 10},
			target: 10,
			want:   3,
		},
		{
			name:   "non empty array with repetition",
			args:   []int{1, 0, 0, 1},
			target: 0,
			want:   2,
		},
		{
			name:   "non empty array with repetition at start",
			args:   []int{1, 1, 0, 0},
			target: 1,
			want:   1,
		},
		{
			name:   "non empty array with repetition at both ends",
			args:   []int{1, 1, 0, 0, 1, 1},
			target: 1,
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LastIndexOf(tt.args, tt.target), "LastIndexOf(%v)", tt.args)
		})
	}
}

func TestPrintAllOccurrence(t *testing.T) {

	tests := []struct {
		name   string
		a      []int
		target int
		wantW  string
	}{
		{
			name:   "empty array",
			a:      []int{},
			target: 2,
			wantW:  "",
		},
		{
			name:   "single element array with target",
			a:      []int{1},
			target: 1,
			wantW:  "0",
		},
		{
			name:   "single element array without target",
			a:      []int{1},
			target: 2,
			wantW:  "",
		},
		{
			name:   "non empty array with repetition",
			a:      []int{1, 2, 3, 2},
			target: 2,
			wantW:  "13",
		},
		{
			name:   "non empty array without repetition",
			a:      []int{1, 2, 3, 4},
			target: 4,
			wantW:  "3",
		},
		{
			name:   "non empty array with all elements repeated",
			a:      []int{1, 1, 1, 1},
			target: 1,
			wantW:  "0123",
		},
		{
			name:   "non empty array with all elements repeated but target not present",
			a:      []int{1, 1, 1, 1},
			target: 2,
			wantW:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			PrintAllOccurrence(w, tt.a, tt.target)
			assert.Equalf(t, tt.wantW, w.String(), "PrintAllOccurrence(%v, %v, %v)", w, tt.a, tt.target)
		})
	}
}
