package main

import (
	"html/template"
	"net/http"
)

type QuantityData struct {
	Quantity string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/MainPage.html",
		"static/html/canvasForExactGraph.html",
		"static/html/dropdownButtonExact.html",
		"static/html/headerMenu.html",
		"static/html/formForQuantity.html")
	quantity := r.FormValue("quantity")
	data := QuantityData{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "mainPage", data)
}

func handleFunc() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", mainPage)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
