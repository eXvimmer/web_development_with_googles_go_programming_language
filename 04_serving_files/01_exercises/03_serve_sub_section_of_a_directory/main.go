package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", dogs)
	http.Handle("/pics/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl.html"))
	tmpl.Execute(w, nil)
}
