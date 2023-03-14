package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

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

	vertex1 := Vertex{1, 100, 200}
	vertex2 := Vertex{2, 200, 300}
	vertex3 := Vertex{3, 400, 100}

	// Create edges
	edge1 := Edge{vertex1, vertex2, 16.5}
	edge2 := Edge{vertex2, vertex3, 21.5}

	// Create graph
	graph := Graph{
		Vertices: []Vertex{vertex1, vertex2, vertex3},
		Edges:    []Edge{edge1, edge2},
	}

	w.Header().Set("Content-Type", "text/html")
	graphJson, err := json.Marshal(graph)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "<script>\n"+
		"var graph = %s;\n"+
		"drawGraph(graph);\n"+
		"</script>\n", graphJson)
	t.ExecuteTemplate(w, "pageForExactGraph", data)
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
