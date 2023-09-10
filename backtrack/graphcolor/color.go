package graphcolor

func colorGraphHelper(g *Graph, colors []int, vertex int) bool {
	if vertex == g.vertices {
		return true // All vertices are colored.
	}

	for color := 1; color <= g.vertices; color++ {
		if g.IsSafe(vertex, color, colors) {
			colors[vertex] = color

			// Recur for the next vertex.
			if colorGraphHelper(g, colors, vertex+1) {
				return true
			}

			// If coloring with the current color doesn't lead to a solution, backtrack.
			colors[vertex] = 0
		}
	}

	return false // No solution exists.
}

func ColorGraph(g *Graph) []int {
	colors := make([]int, g.vertices)
	if !colorGraphHelper(g, colors, 0) {
		return nil
	}
	return colors
}
