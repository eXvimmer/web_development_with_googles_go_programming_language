package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("the / route"))
	})

	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("doggy path"))
	})

	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mustafa Hayati"))
	})
	http.ListenAndServe(":8080", nil)
}
