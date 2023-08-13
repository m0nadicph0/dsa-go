package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Case 0",
			input: []int{3},
			want:  []int{3},
		},
		{
			name:  "Case 1",
			input: []int{3, 2, 1},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Case 2",
			input: []int{3, 2, 1, 4, 6, 5},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "Case 3",
			input: []int{4, 3, 2, 1},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Case 2",
			input: []int{1, 2, 2, 1},
			want:  []int{1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.input)
			assert.Equal(t, tt.want, tt.input, "SelectionSort")
		})
	}
}

func TestMinWithPos(t *testing.T) {

	tests := []struct {
		name         string
		input        []int
		start        int
		wantValue    int
		wantPosition int
	}{
		{
			name:         "Case 1",
			input:        []int{2, 4, 1},
			start:        0,
			wantValue:    1,
			wantPosition: 2,
		},
		{
			name:         "Case 2",
			input:        []int{2, 4, 1, 5, 10, 20},
			start:        1,
			wantValue:    1,
			wantPosition: 2,
		},
		{
			name:         "Case 3",
			input:        []int{1, 2, 3},
			start:        2,
			wantValue:    3,
			wantPosition: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotPosition := minWithPos(tt.input, tt.start)
			assert.Equalf(t, tt.wantValue, gotValue, "minWithPos(%v, %v)", tt.input, tt.start)
			assert.Equalf(t, tt.wantPosition, gotPosition, "minWithPos(%v, %v)", tt.input, tt.start)
		})
	}
}
