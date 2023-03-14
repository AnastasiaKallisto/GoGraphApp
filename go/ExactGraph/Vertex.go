package ExactGraph

type Vertex struct {
	Number int
	X      int
	Y      int
}

func (g *Graph) AddVertex(v Vertex) bool {
	if g.ContainsVertex(v) {
		return false
	}
	g.Vertices = append(g.Vertices, v)
	return true
}

func (g *Graph) ContainsVertex(v Vertex) bool {
	for _, vertex := range g.Vertices {
		if vertex.isEqual(v) {
			return true
		}
	}
	return false
}

func (v1 Vertex) isEqual(v2 Vertex) bool {
	return v1.Number == v2.Number
}
