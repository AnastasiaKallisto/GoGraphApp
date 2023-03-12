package main

import (
	"html/template"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/MainPage.html",
		"static/html/canvasForExactGraph.html",
		"static/html/dropdownButtonExact.html",
		"static/html/headerMenu.html",
		"static/html/textDescription.html")
	t.ExecuteTemplate(w, "mainPage", nil)
}

func handleFunc() {
	http.HandleFunc("/", mainPage)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	handleFunc()
}
