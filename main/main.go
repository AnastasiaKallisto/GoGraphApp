package main

import (
	"html/template"
	"net/http"
)

type QuantityData struct {
	Quantity string
}

func exactGraphPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/exactGraph/pageForExactGraph.html",
		"static/html/exactGraph/canvasForExactGraph.html",
		"static/html/exactGraph/dropdownButtonExact.html",
		"static/html/common/headerMenu.html",
		"static/html/common/formForQuantity.html")
	quantity := r.FormValue("quantity")
	data := QuantityData{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "pageForExactGraph", data)
}

func handleFunc() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/exact", exactGraphPage)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
