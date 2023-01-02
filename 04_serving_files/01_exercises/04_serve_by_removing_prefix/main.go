package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.tmpl.html"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
