package main

import (
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
		"static/html/exactGraph/checkoutIntervalGraphForm.html",
		"static/html/common/quantityForm.html")
	quantity := r.FormValue("quantity")

	data := Data{
		Quantity: quantity,
	}
	t.ExecuteTemplate(w, "pageForExactGraph", data)
}

func switchPageHandler(w http.ResponseWriter, r *http.Request) {
	var redirectURL string
	if r.URL.Path == "/exact" {
		redirectURL = "/interval"
	} else {
		redirectURL = "/exact"
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func intervalGraphPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"static/html/intervalGraph/pageForIntervalGraph.html",
		"static/html/intervalGraph/containerForIntervalGraphs.html",
		"static/html/intervalGraph/dropdownButtonInterval.html",
		"static/html/intervalGraph/checkoutExactGraphForm.html",
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

func handleFunc() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/interval", intervalGraphPage)
	http.HandleFunc("/exact", exactGraphPage)
	http.HandleFunc("/switch", switchPageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
