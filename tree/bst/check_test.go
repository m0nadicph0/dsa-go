package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBST(t *testing.T) {

	tests := []struct {
		name string
		root *Node
		want bool
	}{
		{
			name: "Case 1",
			root: &Node{
				Value: 10,
				Left: &Node{
					Value: 5,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 15,
					Left:  nil,
					Right: nil,
				},
			},
			want: true,
		},
		{
			name: "Case 2",
			root: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 3,
					Left: &Node{
						Value: 25,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsBST(tt.root), "IsBST(%v)", tt.root)
		})
	}
}
