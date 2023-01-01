package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	var method, uri string
	for scanner.Scan() {
		ln := scanner.Text()
		parts := strings.Fields(ln)
		method = parts[0]
		uri = parts[1]
		break // ignore headers and body
	}
	msg := fmt.Sprintf(`
  <h1>Welcome Bro!</h1>
  <h2>Method</h2>
  <p>The method of your request was: %s</p>
  <h2>URI</h2>
  <p>And the URI was: %s</p>
  `, method, uri)
	fmt.Fprintf(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(msg))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	fmt.Fprintf(c, "\r\n")
	fmt.Fprintf(c, msg)
}
