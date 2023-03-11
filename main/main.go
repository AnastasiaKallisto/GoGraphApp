package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/MainPage.html")
	})

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		http.ServeFile(w, r, "html/MainPage.html")
		fmt.Fprintf(w, "Hello %s", name)
	})

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.ListenAndServe("localhost:8080", nil)
}
