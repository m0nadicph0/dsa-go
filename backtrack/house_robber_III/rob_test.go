package house_robber_III

import "testing"

func TestRobIII(t *testing.T) {

	tests := []struct {
		name string
		root *Node
		want int
	}{
		{
			name: "Case 1",
			root: &Node{
				Value: 3,
				Left: &Node{
					Value: 2,
					Left:  nil,
					Right: &Node{Value: 3},
				},
				Right: &Node{
					Value: 3,
					Left:  nil,
					Right: &Node{Value: 1},
				},
			},
			want: 7,
		},
		{
			name: "Case 2",
			root: &Node{
				Value: 3,
				Left: &Node{
					Value: 4,
					Left:  &Node{Value: 1},
					Right: &Node{Value: 3},
				},
				Right: &Node{
					Value: 5,
					Left:  nil,
					Right: &Node{Value: 1},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobIII(tt.root); got != tt.want {
				t.Errorf("RobIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
