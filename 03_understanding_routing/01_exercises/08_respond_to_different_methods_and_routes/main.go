package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var method, uri, msg string
	for scanner.Scan() {
		ln := scanner.Text()
		parts := strings.Fields(ln)
		method = parts[0]
		uri = parts[1]
		break // ignore headers and body
	}

	msg = fmt.Sprintf(`
  <h1>Welcome Bro!</h1>
  <h2>Method</h2>
  <p>The method of your request was: %s</p>
  <h2>URI</h2>
  <p>And the URI was: %s</p>
  `, method, uri)

	msgNotFound := fmt.Sprintf(`
    <h1>Page Not Found</h1>
    <a href="/">Go back</a>
  `)

	switch {
	case method == "GET" && uri == "/":
		respond(c, http.StatusOK, msg, "html")
	case method == "GET" && uri == "/apply":
		msg += `
        <form action="/apply" method="POST">
          <input type="text" name="name" />
          <input type="submit" value="Apply" />
        </form>
      `
		respond(c, http.StatusOK, msg, "html")
	case method == "POST" && uri == "/apply":
		msg += `
        <p>We received your submission</p>
      `
		respond(c, http.StatusOK, msg, "html")
	default:
		respond(c, http.StatusNotFound, msgNotFound, "html")
	}
}

func respond(c net.Conn, status int, msg string, theType string) {
	sl := "HTTP/1.1 " + strconv.Itoa(status) + " " + http.StatusText(status) + "\r\n"
	fmt.Fprintf(c, sl)
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(msg))
	fmt.Fprintf(c, "Content-Type: text/%s\r\n", theType)
	fmt.Fprintf(c, "\r\n")
	fmt.Fprintf(c, msg)
}
