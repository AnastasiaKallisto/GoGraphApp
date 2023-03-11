package ExactGraph

type Graph struct {
	edges    []Edge
	vertices []Vertex
}

func NewGraph() *Graph {
	return &Graph{
		edges:    []Edge{},
		vertices: []Vertex{},
	}
}
