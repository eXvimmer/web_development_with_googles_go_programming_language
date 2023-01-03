package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("times")
	var v int
	if err != nil {
		initilizeVisitTimesCookie(w)
	} else {
		v, err = strconv.Atoi(c.Value)
		if err != nil {
			initilizeVisitTimesCookie(w)
		} else {
			incrementVisitTimesCookie(w, v)
		}
	}

	var msg string
	if v < 1 {
		msg = "Welcome to this page"
	} else {
		msg = fmt.Sprintf("You've visited this page %d times", v+1)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func initilizeVisitTimesCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: "times", Value: "1"})
}

func incrementVisitTimesCookie(w http.ResponseWriter, times int) {
	http.SetCookie(w, &http.Cookie{Name: "times", Value: fmt.Sprintf("%d", times+1)})
}
