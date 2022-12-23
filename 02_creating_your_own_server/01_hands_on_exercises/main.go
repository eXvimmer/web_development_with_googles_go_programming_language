package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// TODO: create a server that returns the URL of the GET request
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		defer conn.Close()
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fields := strings.Fields(ln) // [GET /here/or/there?page=3 HTTP/1.1]
		m := fields[0]
		uri := fields[1]
		if m == "GET" {
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(uri))
			fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, uri)
			conn.Close()
		}
		break
	}
}
