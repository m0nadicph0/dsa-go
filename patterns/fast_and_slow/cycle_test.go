package fast_and_slow

import "testing"

func TestDetectCycle(t *testing.T) {

	tests := []struct {
		name string
		head *Node
		want bool
	}{
		{
			name: "Case 1",
			head: listWithCycle(),
			want: true,
		},
		{
			name: "Case 2",
			head: listWithoutCycle(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycle(tt.head); got != tt.want {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func listWithoutCycle() *Node {
	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	return node1
}

func listWithCycle() *Node {
	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node2
	return node1
}
