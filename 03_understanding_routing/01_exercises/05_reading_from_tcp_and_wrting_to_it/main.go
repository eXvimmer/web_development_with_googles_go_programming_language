package main

import (
	"bufio"
	"fmt"
	// "io"
	"log"
	"net"
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
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		// NOTE: we should break out of this loop at some point;  we break after
		// reading request headers: note, this is not the best solution.
		if text == "" {
			break
		}
	}

	// NOTE: as long as there is a connection, this line won't print; if the
	// client ends the connection, then this line will print
	fmt.Println("code got here")
	msg := "I see you connected"

	// NOTE: wrong way to this this, we should write the response headers and
	// body; use the code below instead
	// io.WriteString(conn, msg)

	// NOTE: the right way to respond
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(msg))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, msg)
}
