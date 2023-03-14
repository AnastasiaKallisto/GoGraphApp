package ExactGraph

type Edge struct {
	A      Vertex  `json:"a"`
	B      Vertex  `json:"b"`
	Weight float64 `json:"weight"`
}

func NewEdge(a Vertex, b Vertex, w float64) *Edge {
	return &Edge{
		A:      a,
		B:      b,
		Weight: w,
	}
}

func (g *Graph) AddEdgeInVertexFormat(a Vertex, b Vertex, weight float64) bool {
	return g.AddEdge(NewEdge(a, b, weight))
}

// AddEdge добавляет ребро в граф, если там уже не было ребра, соединяющего эти вершины
func (g *Graph) AddEdge(e *Edge) bool {
	if g.ContainsEdge(e) {
		return false
	}
	g.Edges = append(g.Edges, *e)
	return true
}

// ContainsEdge проверяет, не содержит ли уже граф ребро, соединяющее те же самые вершины
func (g *Graph) ContainsEdge(e *Edge) bool {
	for _, edge := range g.Edges {
		if edge.isEqual(e) {
			return true
		}
	}
	return false
}

func (e *Edge) isEqual(o *Edge) bool {
	return noOrderEqual([]Vertex{e.A, e.B}, []Vertex{o.A, o.B})
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

func (e *Edge) CompareTo(o *Edge) float64 {
	return e.Weight - o.Weight
}
