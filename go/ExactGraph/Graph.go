package ExactGraph

type Graph struct {
	Edges    []Edge   `json:"edges"`
	Vertices []Vertex `json:"vertices"`
}

func NewGraph() *Graph {
	return &Graph{
		Edges:    []Edge{},
		Vertices: []Vertex{},
	}
}
