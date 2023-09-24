package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTInsert(t *testing.T) {

	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "Case 1",
			values: []int{10, 5, 15, 3, 7, 12, 18},
			want:   []int{3, 5, 7, 10, 12, 15, 18},
		},
		{
			name:   "Case 2",
			values: []int{5, 4, 3, 2, 1},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Case 3",
			values: []int{-5, -4, -3, -2, 1},
			want:   []int{-5, -4, -3, -2, 1},
		},
		{
			name:   "Case 4",
			values: []int{42},
			want:   []int{42},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST()
			for _, value := range tt.values {
				tree.Insert(value)
			}
			got := make([]int, 0)
			tree.InOrder(func(node *Node) {
				got = append(got, node.Value)
			})

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBSTPreOrder(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "Case 1",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Case 2",
			values: []int{5},
			want:   []int{5},
		},
		{
			name:   "Case 3",
			values: []int{5, 4, 3, 2, 1},
			want:   []int{5, 4, 3, 2, 1},
		},
		{
			name:   "Case 4",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Case 5",
			values: []int{10, 5, 15, 3, 7, 12, 18},
			want:   []int{10, 5, 3, 7, 15, 12, 18},
		},
		{
			name:   "Case 6",
			values: []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85},
			want:   []int{50, 30, 20, 10, 25, 40, 35, 45, 70, 60, 55, 65, 80, 75, 85},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST()
			for _, value := range tt.values {
				tree.Insert(value)
			}
			got := make([]int, 0)
			tree.PreOrder(func(node *Node) {
				got = append(got, node.Value)
			})

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBSTPostOrder(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "Case 1",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Case 2",
			values: []int{5},
			want:   []int{5},
		},
		{
			name:   "Case 3",
			values: []int{5, 3, 2, 4, 7, 6, 8},
			want:   []int{2, 4, 3, 6, 8, 7, 5},
		},
		{
			name:   "Case 4",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
		},
		{
			name:   "Case 5",
			values: []int{3, 2, 4, 1, 5},
			want:   []int{1, 2, 5, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST()
			for _, value := range tt.values {
				tree.Insert(value)
			}
			got := make([]int, 0)
			tree.PostOrder(func(node *Node) {
				got = append(got, node.Value)
			})

			assert.Equal(t, tt.want, got)
		})
	}
}
