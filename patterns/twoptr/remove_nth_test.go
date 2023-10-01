package twoptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_RemoveNth(t *testing.T) {

	tests := []struct {
		name    string
		initial []int
		n       int
		want    []int
	}{
		{
			name:    "Case 1",
			initial: []int{1, 2, 3, 4},
			n:       1,
			want:    []int{1, 2, 3},
		},
		{
			name:    "Case 2",
			initial: []int{23, 28, 10, 5, 67, 39, 70, 28},
			n:       2,
			want:    []int{23, 28, 10, 5, 67, 39, 28},
		},
		{
			name:    "Case 3",
			initial: []int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
			n:       3,
			want:    []int{34, 53, 6, 95, 38, 28, 17, 16, 76},
		},
		{
			name:    "Case 4",
			initial: []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			n:       4,
			want:    []int{288, 224, 275, 390, 4, 330, 60, 193},
		},
		{
			name:    "Case 5",
			initial: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			n:       1,
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:    "Case 6",
			initial: []int{69, 8, 49, 106, 116, 112},
			n:       6,
			want:    []int{8, 49, 106, 116, 112},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList()
			for i := len(tt.initial) - 1; i >= 0; i-- {
				list.Insert(tt.initial[i])
			}
			list.RemoveNth(tt.n)
			assert.Equal(t, tt.want, list.ToSlice())
		})
	}
}
