package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

type Vertex struct {
	Number int `json:"number"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

func (g *ExactGraph) AddVertex(v Vertex) bool {
	if g.ContainsVertex(v) {
		return false
	}
	g.Vertices = append(g.Vertices, v)
	return true
}

func (g *ExactGraph) ContainsVertex(v Vertex) bool {
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

func (g *ExactGraph) AddEdgeInVertexFormat(a Vertex, b Vertex, weight float64) bool {
	return g.AddEdge(NewEdge(a, b, weight))
}

// AddEdge добавляет ребро в граф, если там уже не было ребра, соединяющего эти вершины
func (g *ExactGraph) AddEdge(e *Edge) bool {
	if g.ContainsEdge(e) {
		return false
	}
	g.Edges = append(g.Edges, *e)
	return true
}

// ContainsEdge проверяет, не содержит ли уже граф ребро, соединяющее те же самые вершины
func (g *ExactGraph) ContainsEdge(e *Edge) bool {
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

type ExactGraph struct {
	Edges    []Edge   `json:"edges"`
	Vertices []Vertex `json:"vertices"`
}

func NewGraph() *ExactGraph {
	return &ExactGraph{
		Edges:    []Edge{},
		Vertices: []Vertex{},
	}
}

func createRandomGraph(n int) *ExactGraph {
	g := NewGraph()
	if n < 4 {
		return nil
	}
	for i := 0; i < n; i++ {
		g.AddVertex(Vertex{i, 0, 0})
	}

	quantityOfActedVertices := 2
	firstNumber, secondNumber := 0, 0

	g.AddEdgeInVertexFormat(g.Vertices[1], g.Vertices[2], float64(rand.Intn(100)+1))

	for quantityOfActedVertices < n {
		firstNumber = rand.Intn(quantityOfActedVertices) + 1
		quantityOfActedVertices++
		secondNumber = quantityOfActedVertices
		g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
	}

	randomQuantity := rand.Intn(2*n) + 1
	for i := 1; i < randomQuantity; i++ {
		firstNumber = rand.Intn(n) + 1
		secondNumber = rand.Intn(n) + 1
		if firstNumber != secondNumber {
			g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
		}
	}

	return g
}

type Data struct {
	Quantity       string
	SwitchExact    bool
	SwitchInterval bool
}

func exactGraphPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/common/canvasForGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/common/headerMenu.html",
		"static/html/common/clearForm.html",
		"static/html/common/quantityForm.html")
	quantity := r.FormValue("quantity")

	data := Data{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "pageForExactGraph", data)
	fmt.Fprintf(w,
		"</body>\n"+
			"</html>")
}

func intervalGraphPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"static/html/intervalGraph/pageForIntervalGraph.html",
		"static/html/intervalGraph/containerForIntervalGraphs.html",
		"static/html/intervalGraph/dropdownButtonInterval.html",
		"static/html/intervalGraph/formForIntervalGraphInfo.html",
		"static/html/common/headerMenu.html",
		"static/html/common/clearForm.html",
		"static/html/common/quantityForm.html")
	quantity := r.FormValue("quantity")
	data := Data{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "pageForIntervalGraph", data)
}

func drawGraph(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/common/canvasForGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/common/headerMenu.html",
		"static/html/common/clearForm.html",
		"static/html/common/quantityForm.html")
	quantity := r.FormValue("quantity")
	data := Data{
		Quantity: quantity,
	}

	vertex1 := Vertex{1, 100, 200}
	vertex2 := Vertex{2, 200, 300}
	vertex3 := Vertex{3, 400, 100}

	// Create edges
	edge1 := Edge{vertex1, vertex2, 16.5}
	edge2 := Edge{vertex2, vertex3, 21.5}

	// Create graph
	graph := ExactGraph{
		Vertices: []Vertex{vertex1, vertex2, vertex3},
		Edges:    []Edge{edge1, edge2},
	}

	graphJson, err := json.Marshal(graph)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "pageForExactGraph", data)
	fmt.Fprintf(w, "<script>\n"+
		"var graph = %s;\n"+
		"drawGraph(graph);\n"+
		"</script>\n"+
		"</body>\n"+
		"</html>", graphJson)
}

func handleFunc() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/interval", intervalGraphPage)
	http.HandleFunc("/exact", exactGraphPage)
	http.HandleFunc("/exact/draw", drawGraph)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
