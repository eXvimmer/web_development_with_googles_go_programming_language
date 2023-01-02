package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/puppies.webp", puppies)
	http.HandleFunc("/dog/", dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo ran"))
}

func dog(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dog.tmpl.html"))
	tmpl.Execute(w, nil)
}

func puppies(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "puppies.webp")
}
