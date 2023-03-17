package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var graph *ExactGraph
var typeOfGraph string   // "exact" "interval"
var sourceOfGraph string // "generate" "fromFile"
var quantity string      // число
var MSTPrimExact *ExactGraph
var MSTKruscalExact *ExactGraph

//var arrayMSTPrimInterval []*IntervalGraph
//var arrayMSTKruscalInterval []*IntervalGraph

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
	if g.GetEqualEdge(e) != nil {
		return false
	}
	g.Edges = append(g.Edges, *e)
	return true
}

// GetEqualEdge возвращает ребро, вершины которого совпадают с тем, что пришло в функцию
func (g *ExactGraph) GetEqualEdge(e *ExactEdge) *ExactEdge {
	for _, edge := range g.Edges {
		if edge.isEqual(e) {
			return &edge
		}
	}
	return nil
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

func (g *ExactGraph) Prim() *ExactGraph {
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
	fmt.Println("Массив вершин графа: ", g.Vertices)
	minEdge := ExactEdge{
		Weight: math.MaxFloat64,
	}
	// Пока не посетим все вершины
	for len(visited) < len(g.Vertices) {
		// Будем искать ребро минимального веса, соединяющее посещенную и непосещенную вершину
		// Установим максимальный вес, а с вершинами разберемся позже
		minEdge.Weight = math.MaxFloat64
		fmt.Println("Массив посещенных вершин: ", visited)
		fmt.Println("Массив вершин дерева mst: ", MSTPrimExact.Vertices)
		// пройдем по всем посещенным вершинам
		for _, v1 := range MSTPrimExact.Vertices {
			// пройдем по всем вершинам графа
			for _, v2 := range g.Vertices {
				// но смотрим только на те, которые непосещенные
				if visited[v2.Number] {
					fmt.Println("Вершина уже посещенная: ", v2.Number)
					continue
				}
				fmt.Println("Вершина НЕ посещенная: ", v2.Number)
				// кандидат на новое ребро в мин. ост. дереве
				potentialEdge := ExactEdge{
					A: v1,
					B: v2,
				}
				fmt.Println("Потенциальное ребро: ", potentialEdge)
				// а содержится ли оно в графе?
				graphEdge := g.GetEqualEdge(&potentialEdge)
				fmt.Println("Аналогичное ребро из графа: ", graphEdge)
				if graphEdge == nil {
					continue
				}
				// Может, у него вес меньше? Если да, то оно лучше
				if graphEdge.Weight < minEdge.Weight {
					minEdge = *graphEdge
				}
			}
		}
		// В конце цикла нашли нужное ребро с минимальным весом
		// добавляем его
		fmt.Println("В итоге добавляем minEdge &minEdge: ", minEdge, &minEdge)
		MSTPrimExact.AddEdge(&minEdge)
		// если A была посещенной, значит, надо добавить B, иначе А
		if visited[minEdge.A.Number] {
			visited[minEdge.B.Number] = true
			MSTPrimExact.Vertices = append(MSTPrimExact.Vertices, minEdge.B)
			fmt.Println("В граф добавили вершину: ", minEdge.B)
		} else {
			visited[minEdge.A.Number] = true
			MSTPrimExact.Vertices = append(MSTPrimExact.Vertices, minEdge.A)
			fmt.Println("В граф добавили вершину: ", minEdge.A)
		}
	}
	fmt.Println("Минимальное ост дерево: ", MSTPrimExact)
	// на выходе имеем мин. ост. дерево
	return MSTPrimExact
}

func (g *ExactGraph) Kruskal() *ExactGraph {
	// Создаём новый граф только с вершинами
	resultGraph := NewGraph()
	// Копируем все рёбра и сортируем их по весу
	edges := g.Edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	// Алгоритм Краскала
	for _, e := range edges {
		// Если добавление ребра приведёт к циклу - не добавляем
		if !searchChain(e.A, e.B, resultGraph.Edges, []Vertex{}) {
			resultGraph.AddEdge(&e)
			resultGraph.AddVertex(e.A)
			resultGraph.AddVertex(e.B)
		}
	}
	return resultGraph
}

func searchChain(a Vertex, b Vertex, graphEdges []ExactEdge, checkedVertices []Vertex) bool {
	var verticesThatNeedToBeChecked []Vertex
	for _, edge := range graphEdges {
		if edge.A == a {
			if edge.B == b {
				return true
			}
			if !containsVertex(checkedVertices, edge.B) {
				verticesThatNeedToBeChecked = append(verticesThatNeedToBeChecked, edge.B)
			}
		}
		if edge.B == a {
			if edge.A == b {
				return true
			}
			if !containsVertex(checkedVertices, edge.A) {
				verticesThatNeedToBeChecked = append(verticesThatNeedToBeChecked, edge.A)
			}
		}
	}
	checkedVertices = append(checkedVertices, a)
	for _, vertex := range verticesThatNeedToBeChecked {
		if searchChain(vertex, b, graphEdges, checkedVertices) {
			return true
		}
	}
	return false
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
	MSTKruscalExact = nil
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
	prim := graph.Prim()
	fmt.Println("Посчитали прима ")
	primJson, err := json.Marshal(prim)
	if err != nil {
		panic(err)
	}
	kruscal := graph.Kruskal()
	fmt.Println("Посчитали краскала ")
	KruscalJson, err := json.Marshal(kruscal)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "pageForExactGraph", data)
	fmt.Println("Прогрузил страницу")
	fmt.Fprintf(w, "<script>\n"+
		"var graph = %s;\n"+
		"drawGraph(graph);\n", graphJson)
	fmt.Fprintf(w,
		"var primGraph = %s;\n"+
			"drawPrimTree(primGraph);\n", primJson)
	fmt.Fprintf(w,
		"var KruscalGraph = %s;\n"+
			"drawCruscalTree(KruscalGraph);\n", KruscalJson)
	fmt.Println("Загрузил скрипт")
	fmt.Fprintf(w,
		"document.getElementById('textGraph').value = '%s';\n"+
			"document.getElementById('MSTPrim').value = '%s';\n"+
			"document.getElementById('MSTCruscal').value = '%s';\n"+
			"</script>\n"+
			"</body>\n"+
			"</html>", graphToReadingFormat(graph), graphToReadingFormat(prim), graphToReadingFormat(kruscal))
	fmt.Println("присвоил значения текстовые")
}

func graphToReadingFormat(graph *ExactGraph) string {
	answer := "Ребра:\\n"
	for _, edge := range graph.Edges {
		answer = strings.Join([]string{
			answer,
			"V",
			strconv.Itoa(edge.A.Number),
			" V",
			strconv.Itoa(edge.B.Number),
			" W = ",
			fmt.Sprintf("%.2f", edge.Weight),
			"\\n",
		}, "")
	}
	return answer
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
