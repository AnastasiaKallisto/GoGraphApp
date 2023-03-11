package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/MainPage.html")

	})

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		http.ServeFile(w, r, "static/MainPage.html")
		fmt.Fprintf(w, "Hello %s", name)
	})
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe("localhost:8080", nil)
}
