package ExactGraph

type Edge struct {
	vertices []Vertex
	weight   int
}

func NewEdge(a Vertex, b Vertex, w int) *Edge {
	return &Edge{
		vertices: []Vertex{a, b},
		weight:   w,
	}
}

func (g *Graph) AddEdgeInVertexFormat(a Vertex, b Vertex, weight int) bool {
	return g.AddEdge(NewEdge(a, b, weight))
}

// AddEdge добавляет ребро в граф, если там уже не было ребра, соединяющего эти вершины
func (g *Graph) AddEdge(e *Edge) bool {
	if g.ContainsEdge(e) {
		return false
	}
	g.edges = append(g.edges, *e)
	return true
}

// ContainsEdge проверяет, не содержит ли уже граф ребро, соединяющее те же самые вершины
func (g *Graph) ContainsEdge(e *Edge) bool {
	for _, edge := range g.edges {
		if edge.isEqual(e) {
			return true
		}
	}
	return false
}

func (e *Edge) isEqual(o *Edge) bool {
	return noOrderEqual(e.vertices, o.vertices)
}

func noOrderEqual(a []Vertex, b []Vertex) bool {
	if len(a) != len(b) {
		return false
	}
	visited := map[int]bool{}
	for _, valueA := range a {
		found := false
		for j, valueB := range b {
			if !visited[j] && valueA.Number == valueB.Number {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (e *Edge) CompareTo(o *Edge) int {
	return e.weight - o.weight
}
