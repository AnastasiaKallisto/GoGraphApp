package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var graph *ExactGraph
var typeOfGraph string   // "exact" "interval"
var sourceOfGraph string // "generate" "fromFile"
var quantity string      // число
var MSTPrimExact *ExactGraph
var MSTCruscalExact *ExactGraph

//var arrayMSTPrimInterval []*IntervalGraph
//var arrayMSTCruscalInterval []*IntervalGraph

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

type ExactEdge struct {
	A      Vertex  `json:"a"`
	B      Vertex  `json:"b"`
	Weight float64 `json:"weight"`
}

func NewEdge(a Vertex, b Vertex, w float64) *ExactEdge {
	return &ExactEdge{
		A:      a,
		B:      b,
		Weight: w,
	}
}

func (g *ExactGraph) AddEdgeInVertexFormat(a Vertex, b Vertex, weight float64) bool {
	return g.AddEdge(NewEdge(a, b, weight))
}

// AddEdge добавляет ребро в граф, если там уже не было ребра, соединяющего эти вершины
func (g *ExactGraph) AddEdge(e *ExactEdge) bool {
	if g.ContainsEdge(e) {
		return false
	}
	g.Edges = append(g.Edges, *e)
	return true
}

// ContainsEdge проверяет, не содержит ли уже граф ребро, соединяющее те же самые вершины
func (g *ExactGraph) ContainsEdge(e *ExactEdge) bool {
	for _, edge := range g.Edges {
		if edge.isEqual(e) {
			return true
		}
	}
	return false
}

func (e *ExactEdge) isEqual(o *ExactEdge) bool {
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

func (e *ExactEdge) CompareTo(o *ExactEdge) float64 {
	return e.Weight - o.Weight
}

type ExactGraph struct {
	Edges    []ExactEdge `json:"edges"`
	Vertices []Vertex    `json:"vertices"`
}

func NewGraph() *ExactGraph {
	return &ExactGraph{
		Edges:    []ExactEdge{},
		Vertices: []Vertex{},
	}
}

func createGraphFromFile(content string) (*ExactGraph, error) {
	graph := NewGraph()
	lines := strings.Split(content, "\n")
	quantity, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		return nil, fmt.Errorf("invalid content format: %v", err)
	}
	if quantity > len(lines)-2 {
		return nil, fmt.Errorf("invalid content format: quantity of edges must me >= %v", quantity)
	}
	for i := 1; i < len(lines); i++ {
		data := strings.Split(lines[i], " ")
		if len(data) != 3 {
			return nil, fmt.Errorf("invalid content format: in line must be number1, number2 and weight, not this %v", data)
		}
		number1, err := strconv.Atoi(strings.TrimSpace(data[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid content format: %v", err)
		}
		number2, err := strconv.Atoi(strings.TrimSpace(data[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid content format: %v", err)
		}
		weight, err := strconv.ParseFloat(strings.TrimSpace(data[2]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid content format: %v", err)
		}
		vertex1 := Vertex{
			Number: number1,
		}
		vertex2 := Vertex{
			Number: number2,
		}
		graph.AddVertex(vertex1)
		graph.AddVertex(vertex2)
		graph.AddEdgeInVertexFormat(vertex1, vertex2, weight)
	}
	return graph, nil
}

func createRandomGraph(n int) *ExactGraph {
	g := NewGraph()
	if n < 3 {
		return nil
	}
	for i := 0; i < n; i++ {
		g.AddVertex(Vertex{i, 0, 0})
	}

	quantityOfActedVertices := 2
	firstNumber, secondNumber := 0, 0

	actedNumbersOfVertices := []int{0, 1}
	g.AddEdgeInVertexFormat(g.Vertices[0], g.Vertices[1], float64(rand.Intn(100)+1))

	for quantityOfActedVertices < n {
		firstNumber = actedNumbersOfVertices[rand.Intn(quantityOfActedVertices)]
		secondNumber = quantityOfActedVertices
		quantityOfActedVertices++
		actedNumbersOfVertices = append(actedNumbersOfVertices, secondNumber)
		g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
	}

	randomQuantity := rand.Intn(2*n) + 1
	for i := 1; i < randomQuantity; i++ {
		firstNumber = rand.Intn(n)
		secondNumber = rand.Intn(n)
		if firstNumber != secondNumber {
			g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
		}
	}

	return g
}

func containsVertex(vertices []Vertex, a Vertex) bool {
	for _, vertex := range vertices {
		if vertex == a {
			return true
		}
	}
	return false
}

func Prim(g ExactGraph) *ExactGraph {

	// будущее минимальное остовное дерево
	MSTPrimExact = NewGraph()

	// создадим отображение номеров посещенных вершин к факту их посещения
	visited := make(map[int]bool)
	// Первая вершина с индексом 0
	startingVertex := g.Vertices[0]
	// Теперь она посещенная
	visited[startingVertex.Number] = true
	// и она есть в списке вершин минимального остовного дерева
	MSTPrimExact.Vertices = append(MSTPrimExact.Vertices, startingVertex)

	// Пока не посетим все вершины
	for len(visited) < len(g.Vertices) {
		// Будем искать ребро минимального веса, соединяющее посещенную и непосещенную вершину
		// Установим максимальный вес, а с вершинами разберемся позже
		minEdge := ExactEdge{
			Weight: math.MaxFloat64,
		}
		// пройдем по всем посещенным вершинам
		for _, v1 := range MSTPrimExact.Vertices {
			// пройдем по всем вершинам графа
			for _, v2 := range g.Vertices {
				// но смотрим только на те, которые непосещенные
				if visited[v2.Number] {
					continue
				}
				// кандидат на новое ребро в мин. ост. дереве
				potentialEdge := ExactEdge{
					A:      v1,
					B:      v2,
					Weight: math.Sqrt(float64(squareDistance(v1, v2))),
				}
				// Может, у него вес меньше? Если да, то оно лучше
				if potentialEdge.Weight < minEdge.Weight {
					minEdge = potentialEdge
				}
			}
		}

		// В конце цикла нашли нужное ребро с минимальным весом
		// добавляем его
		MSTPrimExact.AddEdge(&minEdge)
		// если A была посещенной, значит, надо добавить B, иначе А
		if visited[minEdge.A.Number] {
			visited[minEdge.B.Number] = true
			MSTPrimExact.Vertices = append(MSTPrimExact.Vertices, minEdge.B)
		} else {
			visited[minEdge.A.Number] = true
			MSTPrimExact.Vertices = append(MSTPrimExact.Vertices, minEdge.A)
		}
	}
	// на выходе имеем мин. ост. дерево
	return MSTPrimExact
}

func squareDistance(v1, v2 Vertex) int {
	return ((v1.X - v2.X) * (v1.X - v2.X)) + ((v1.Y - v2.Y) * (v1.Y - v2.Y))
}

type Data struct {
	Quantity string
	FileName string
}

func exactGraphPage(w http.ResponseWriter, r *http.Request) {
	graph = nil
	sourceOfGraph = ""
	quantity = ""
	MSTPrimExact = nil
	MSTCruscalExact = nil
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/common/canvasForGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/exactGraph/uploadExactFileForm.html",
		"static/html/common/headerMenu.html",
		"static/html/exactGraph/exactClearForm.html",
		"static/html/exactGraph/exactQuantityForm.html",
	)
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
		"static/html/intervalGraph/intervalClearForm.html",
		"static/html/intervalGraph/intervalQuantityForm.html",
		"static/html/intervalGraph/uploadIntervalFileForm.html")
	quantity := r.FormValue("quantity")
	data := Data{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "pageForIntervalGraph", data)
}

func generateExactGraphPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/common/canvasForGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/exactGraph/uploadExactFileForm.html",
		"static/html/common/headerMenu.html",
		"static/html/exactGraph/exactClearForm.html",
		"static/html/exactGraph/exactQuantityForm.html")
	data := Data{}
	fmt.Println(quantity)
	if quantity != "" {
		data = Data{
			Quantity: quantity,
		}
	} else {
		quantity = r.FormValue("quantity")
		data = Data{
			Quantity: quantity,
		}
		n, _ := strconv.Atoi(quantity)
		if graph == nil {
			graph = createRandomGraph(n)
			sourceOfGraph = "generate"
		}
	}
	graphJson, err := json.Marshal(graph)
	if err != nil {
		panic(err)
	}
	primJson, err := json.Marshal(Prim(*graph))
	if err != nil {
		panic(err)
	}
	cruscalJson, err := json.Marshal(Prim(*graph))
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "pageForExactGraph", data)
	fmt.Fprintf(w, "<script>\n"+
		"var graph = %s;\n"+
		"drawGraph(graph);\n", graphJson)
	fmt.Fprintf(w,
		"var primGraph = %s;\n"+
			"drawPrimTree(primGraph);\n", primJson)
	fmt.Fprintf(w,
		"var cruscalGraph = %s;\n"+
			"drawCruscalTree(cruscalGraph);\n", cruscalJson)
	fmt.Fprintf(w,
		"document.getElementById('textGraph').value = '%s';\n"+
			"document.getElementById('MSTPrim').value = '%s';\n"+
			"document.getElementById('MSTCruscal').value = '%s';\n"+
			"</script>\n"+
			"</body>\n"+
			"</html>", graphJson, primJson, cruscalJson)
}

func getExactGraphFromFilePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/common/canvasForGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/exactGraph/uploadExactFileForm.html",
		"static/html/common/headerMenu.html",
		"static/html/exactGraph/exactClearForm.html",
		"static/html/exactGraph/exactQuantityForm.html")
	file, _, err := r.FormFile("exactGraphTxtFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if graph == nil {
		graph, err = createGraphFromFile(string(content))
	}

	graphJson, err := json.Marshal(graph)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "pageForExactGraph", nil)
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
	http.HandleFunc("/exact/generate", generateExactGraphPage)
	http.HandleFunc("/exact/from-file", getExactGraphFromFilePage)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
