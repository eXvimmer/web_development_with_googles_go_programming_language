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
	var v int
	c, err := r.Cookie("times")
	if err != nil {
		c = initilizeVisitTimesCookie(w)
	} else {
		v, err = strconv.Atoi(c.Value)
		if err != nil {
			c = initilizeVisitTimesCookie(w)
		} else {
			c = incrementVisitTimesCookie(w, v)
		}
	}

	var msg string
	if v < 1 {
		msg = "Welcome to this page"
	} else {
		msg = fmt.Sprintf("You've visited this page %d times", v+1)
	}
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func initilizeVisitTimesCookie(w http.ResponseWriter) *http.Cookie {
	return &http.Cookie{Name: "times", Value: "1"}
}

func incrementVisitTimesCookie(w http.ResponseWriter, times int) *http.Cookie {
	return &http.Cookie{Name: "times", Value: fmt.Sprintf("%d", times+1)}
}
