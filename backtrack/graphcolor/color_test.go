package graphcolor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Edge struct {
	U int
	V int
}

func TestColorGraph(t *testing.T) {

	tests := []struct {
		name     string
		vertices int
		edges    []Edge
		want     []int
	}{
		{
			name:     "Case 1",
			vertices: 4,
			edges:    []Edge{{0, 1}, {0, 2}, {1, 2}, {2, 3}},
			want:     []int{1, 2, 3, 1},
		},
		{
			name:     "Case 2",
			vertices: 4,
			edges:    []Edge{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			want:     []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := buildGraph(tt.vertices, tt.edges)
			got := ColorGraph(graph)
			assert.Equal(t, tt.want, got)
		})
	}
}

func buildGraph(vertices int, edges []Edge) *Graph {
	graph := NewGraph(vertices)
	for _, edge := range edges {
		graph.AddEdge(edge.U, edge.V)
	}
	return graph
}
