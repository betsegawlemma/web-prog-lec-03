package main

import (
	"html/template"
	"net/http"
)

// Course struct
type Course struct {
	Title, Code, Description string
	ECTS                     int
}

var templ = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	crs := Course{"DLD", "ITSE-3182", "Lorem Ipsum", 7}
	templ.Execute(w, crs)
}

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", mux)
}
