package graphcolor

// Graph Adjacency matrix representation.
type Graph struct {
	vertices int
	matrix   [][]bool
}

func NewGraph(vertices int) *Graph {
	matrix := make([][]bool, vertices)
	for i := range matrix {
		matrix[i] = make([]bool, vertices)
	}
	return &Graph{vertices: vertices, matrix: matrix}
}

func (g *Graph) AddEdge(u, v int) {
	g.matrix[u][v] = true
	g.matrix[v][u] = true
}

func (g *Graph) IsSafe(vertex, color int, colors []int) bool {
	for i := 0; i < g.vertices; i++ {
		if g.matrix[vertex][i] && colors[i] == color {
			return false
		}
	}
	return true
}
