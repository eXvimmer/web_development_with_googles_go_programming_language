package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.tmpl.html", c.Value)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil { // no cookies
		c = &http.Cookie{Name: "session", Value: uuid.New().String()}
		http.SetCookie(w, c)
	}
	return c
}
