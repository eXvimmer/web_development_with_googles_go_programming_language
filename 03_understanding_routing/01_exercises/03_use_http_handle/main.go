package main

import (
	"html/template"
	"net/http"
)

type Data struct {
	Msg   string
	Extra string
}

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.tmpl.html"))
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, Data{Msg: "Welcome to / route"})
}

func dog(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, Data{Msg: "Welcome to /dog route"})
}

func me(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, Data{Msg: "Welcome to " + r.URL.Path + " route",
		Extra: "My name is Mustafa Hayati"})
}
