package fast_and_slow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMiddleNode(t *testing.T) {

	tests := []struct {
		name string
		head *Node
		want int
	}{
		{
			name: "Case 1",
			head: &Node{
				value: 1,
				next: &Node{
					value: 2,
					next: &Node{
						value: 3,
						next: &Node{
							value: 4,
							next: &Node{
								value: 5,
								next:  nil,
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "Case 2",
			head: &Node{
				value: 1,
				next: &Node{
					value: 2,
					next: &Node{
						value: 3,
						next: &Node{
							value: 4,
							next:  nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMiddleNode(tt.head)
			assert.Equal(t, tt.want, got.value)
		})
	}
}
